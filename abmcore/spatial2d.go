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

// Modes are different preset modes for a [Spatial2D] plot.
type Modes int32 //enums:enum -trim-prefix Mode

const (

	// ModeSpatial shows agents based on their spatial positions.
	ModeSpatial Modes = iota

	// ModeBelief shows agents based on their political beliefs.
	ModeBelief
)

// Spatial2D is a 2d plot of a simulation based on the [abm.AgentBase.Position].
type Spatial2D struct {
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

func (sp *Spatial2D) Init() {
	sp.Frame.Init()
	sp.Styler(func(s *styles.Style) {
		s.Grow.Set(1, 1)
		s.Direction = styles.Column
	})

	tree.AddChild(sp, func(w *core.Toolbar) {
		w.Maker(sp.MakeToolbar)
	})
	tree.AddChild(sp, func(w *plotcore.Editor) {
		sp.plot = w

		w.Updater(func() {
			if sp.table == nil {
				sp.MakeTable()
			}
			sp.UpdateTable()
		})
	})
}

// MakeTable creates the data table for plotting.
func (sp *Spatial2D) MakeTable() {
	sp.table = table.New()
	n := len(sp.Sim.Base().Agents)
	sp.table.AddColumn("Spatial X", tensor.NewFloat32(n))
	sp.table.AddColumn("Spatial Y", tensor.NewFloat32(n))
	sp.table.AddColumn("Belief X", tensor.NewFloat32(n))
	sp.table.AddColumn("Belief Y", tensor.NewFloat32(n))

	plot.Styler(sp.table.Column("Spatial Y"), sp.colorStyler)
	plot.Styler(sp.table.Column("Belief Y"), sp.colorStyler)

	plot.Styler(sp.table.Column("Spatial X"), func(s *plot.Style) {
		s.Role = plot.X
	})

	sp.plot.SetTable(sp.table)
}

// UpdateTable updates the data table with the current agent data.
func (sp *Spatial2D) UpdateTable() {
	agents := sp.Sim.Base().Agents
	sp.table.SetNumRows(len(agents))

	for i, a := range agents {
		pos := a.Base().Position
		sp.table.Column("Spatial X").SetFloat(float64(pos.X), i)
		sp.table.Column("Spatial Y").SetFloat(float64(pos.Y), i)

		beliefs := a.Base().Beliefs
		sp.table.Column("Belief X").SetFloat(float64(beliefs[0]), i)
		sp.table.Column("Belief Y").SetFloat(float64(beliefs[1]), i)
	}
}

// colorStyler is a plot styler that styles points based on agent beliefs.
func (sp *Spatial2D) colorStyler(s *plot.Style) {
	s.Line.On = plot.Off
	s.Point.On = plot.On

	agents := sp.Sim.Base().Agents
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

func (sp *Spatial2D) MakeToolbar(p *tree.Plan) {
	tree.Add(p, func(w *core.Switches) {
		core.Bind(&sp.Mode, w)
	})

	sp.plot.MakeToolbar(p)
}
