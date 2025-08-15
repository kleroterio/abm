// Copyright (c) 2025, Kleroterio. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package abmcore

import (
	"image"

	"cogentcore.org/core/colors"
	"cogentcore.org/core/colors/cam/hct"
	"cogentcore.org/core/core"
	"cogentcore.org/core/styles"
	"cogentcore.org/core/tree"
	"cogentcore.org/lab/plot"
	"cogentcore.org/lab/plotcore"
	"cogentcore.org/lab/table"
	"cogentcore.org/lab/tensor"
	"github.com/kleroterio/abm/abm"
)

// Modes are different preset modes for a [Agents] plot.
type Modes int32 //enums:enum -trim-prefix Mode

const (

	// ModeSpatial shows agents based on their spatial positions.
	ModeSpatial Modes = iota

	// ModeBelief shows agents based on their political beliefs.
	ModeBelief
)

// Agents is a customizable 2D plot of the agents in a simulation.
type Agents struct {
	core.Frame

	// Sim is the simulation that this 2D representation is based on.
	Sim abm.Sim

	// Mode is the current preset plotting mode.
	Mode Modes

	// table is the data table for plotting.
	table *table.Table

	// plot is the plot editor widget.
	plot *plotcore.Editor
}

func (ag *Agents) Init() {
	ag.Frame.Init()
	ag.Styler(func(s *styles.Style) {
		s.Grow.Set(1, 1)
		s.Direction = styles.Column
	})

	tree.AddChild(ag, func(w *core.Toolbar) {
		w.Maker(ag.MakeToolbar)
	})
	tree.AddChild(ag, func(w *plotcore.Editor) {
		ag.plot = w

		w.Updater(ag.UpdateTable)
	})
}

// makeTable creates the data table for plotting.
func (ag *Agents) makeTable() {
	ag.table = table.New()
	n := len(ag.Sim.Base().Agents)
	ag.table.AddColumn("Spatial X", tensor.NewFloat32(n))
	ag.table.AddColumn("Spatial Y", tensor.NewFloat32(n))
	ag.table.AddColumn("Belief X", tensor.NewFloat32(n))
	ag.table.AddColumn("Belief Y", tensor.NewFloat32(n))
	ag.table.AddColumn("Influence", tensor.NewFloat32(n))

	plot.Styler(ag.table.Column("Spatial Y"), ag.colorStyler)
	plot.Styler(ag.table.Column("Belief Y"), ag.colorStyler)

	plot.Styler(ag.table.Column("Influence"), func(s *plot.Style) {
		s.Role = plot.Size
	})

	plot.Styler(ag.table.Column("Spatial X"), func(s *plot.Style) {
		if ag.Mode == ModeSpatial {
			s.Role = plot.X
		} else {
			s.Role = plot.Y
		}
	})
	plot.Styler(ag.table.Column("Spatial Y"), func(s *plot.Style) {
		s.On = ag.Mode == ModeSpatial
	})
	plot.Styler(ag.table.Column("Belief X"), func(s *plot.Style) {
		if ag.Mode == ModeBelief {
			s.Role = plot.X
		} else {
			s.Role = plot.Y
		}
	})
	plot.Styler(ag.table.Column("Belief Y"), func(s *plot.Style) {
		s.On = ag.Mode == ModeBelief
	})

	ag.plot.SetTable(ag.table)
}

// UpdateTable updates the data table with the current agent data.
func (ag *Agents) UpdateTable() {
	if ag.table == nil {
		ag.makeTable()
	}

	agents := ag.Sim.Base().Agents
	ag.table.SetNumRows(len(agents))

	for i, a := range agents {
		pos := a.Base().Position
		ag.table.Column("Spatial X").SetFloat(float64(pos.X), i)
		ag.table.Column("Spatial Y").SetFloat(float64(pos.Y), i)

		beliefs := a.Base().Beliefs
		ag.table.Column("Belief X").SetFloat(float64(beliefs[0]), i)
		ag.table.Column("Belief Y").SetFloat(float64(beliefs[1]), i)

		ag.table.Column("Influence").SetFloat(float64(a.Base().Influence), i)
	}
}

// UpdatePlot updates the table and plot.
func (ag *Agents) UpdatePlot() {
	ag.UpdateTable()
	if ag.plot.IsVisible() {
		ag.plot.UpdatePlot()
	}
}

// colorStyler is a plot styler that styles points based on agent beliefs.
func (ag *Agents) colorStyler(s *plot.Style) {
	s.Line.On = plot.Off
	s.Point.On = plot.On
	// need a little extra room to avoid plot shifting
	s.Range.SetMin(-0.02).SetMax(1.02)
	s.Plot.XAxis.Range.SetMin(-0.02).SetMax(1.02)
	s.Point.Size.Pt(5)

	agents := ag.Sim.Base().Agents
	s.Point.ColorFunc = func(i int) image.Image {
		beliefs := agents[i].Base().Beliefs

		hue := 270 + beliefs[0]*120 // blue to red
		chroma, tone := float32(100), float32(50)
		if len(beliefs) >= 2 {
			tone = 25 + beliefs[1]*50
		}
		if len(beliefs) >= 3 {
			chroma = 100 * beliefs[2]
		}

		c := hct.New(hue, chroma, tone)
		return colors.Uniform(c.AsRGBA())
	}
	s.Point.FillFunc = s.Point.ColorFunc
}

func (ag *Agents) MakeToolbar(p *tree.Plan) {
	tree.Add(p, func(w *core.Switches) {
		core.Bind(&ag.Mode, w)
	})

	ag.plot.MakeToolbar(p)
}
