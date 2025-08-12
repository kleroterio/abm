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

// Modes are different preset modes for a [Plot] plot.
type Modes int32 //enums:enum -trim-prefix Mode

const (

	// ModeSpatial shows agents based on their spatial positions.
	ModeSpatial Modes = iota

	// ModeBelief shows agents based on their political beliefs.
	ModeBelief
)

// Plot is a customizable 2D plot of a simulation.
type Plot struct {
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

func (pl *Plot) Init() {
	pl.Frame.Init()
	pl.Styler(func(s *styles.Style) {
		s.Grow.Set(1, 1)
		s.Direction = styles.Column
	})

	tree.AddChild(pl, func(w *core.Toolbar) {
		w.Maker(pl.MakeToolbar)
	})
	tree.AddChild(pl, func(w *plotcore.Editor) {
		pl.plot = w

		w.Updater(func() {
			if pl.table == nil {
				pl.MakeTable()
			}
			pl.UpdateTable()
		})
	})
}

// MakeTable creates the data table for plotting.
func (pl *Plot) MakeTable() {
	pl.table = table.New()
	n := len(pl.Sim.Base().Agents)
	pl.table.AddColumn("Spatial X", tensor.NewFloat32(n))
	pl.table.AddColumn("Spatial Y", tensor.NewFloat32(n))
	pl.table.AddColumn("Belief X", tensor.NewFloat32(n))
	pl.table.AddColumn("Belief Y", tensor.NewFloat32(n))

	plot.Styler(pl.table.Column("Spatial Y"), pl.colorStyler)
	plot.Styler(pl.table.Column("Belief Y"), pl.colorStyler)

	plot.Styler(pl.table.Column("Spatial X"), func(s *plot.Style) {
		if pl.Mode == ModeSpatial {
			s.Role = plot.X
		} else {
			s.Role = plot.Y
		}
	})
	plot.Styler(pl.table.Column("Spatial Y"), func(s *plot.Style) {
		s.On = pl.Mode == ModeSpatial
	})
	plot.Styler(pl.table.Column("Belief X"), func(s *plot.Style) {
		if pl.Mode == ModeBelief {
			s.Role = plot.X
		} else {
			s.Role = plot.Y
		}
	})
	plot.Styler(pl.table.Column("Belief Y"), func(s *plot.Style) {
		s.On = pl.Mode == ModeBelief
	})

	pl.plot.SetTable(pl.table)
}

// UpdateTable updates the data table with the current agent data.
func (pl *Plot) UpdateTable() {
	agents := pl.Sim.Base().Agents
	pl.table.SetNumRows(len(agents))

	for i, a := range agents {
		pos := a.Base().Position
		pl.table.Column("Spatial X").SetFloat(float64(pos.X), i)
		pl.table.Column("Spatial Y").SetFloat(float64(pos.Y), i)

		beliefs := a.Base().Beliefs
		pl.table.Column("Belief X").SetFloat(float64(beliefs[0]), i)
		pl.table.Column("Belief Y").SetFloat(float64(beliefs[1]), i)
	}
}

// colorStyler is a plot styler that styles points based on agent beliefs.
func (pl *Plot) colorStyler(s *plot.Style) {
	s.Line.On = plot.Off
	s.Point.On = plot.On

	agents := pl.Sim.Base().Agents
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

func (pl *Plot) MakeToolbar(p *tree.Plan) {
	tree.Add(p, func(w *core.Switches) {
		core.Bind(&pl.Mode, w)
	})

	pl.plot.MakeToolbar(p)
}
