// Copyright (c) 2025, Kleroterio. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package abmcore

import (
	"image"

	"cogentcore.org/core/colors"
	"cogentcore.org/core/colors/cam/hct"
	"cogentcore.org/core/styles"
	"cogentcore.org/lab/plot"
	"cogentcore.org/lab/plot/plots"
	"cogentcore.org/lab/plotcore"
	"cogentcore.org/lab/tensor"
	"github.com/kleroterio/abm/abm"
)

// Spatial2D is a 2d plot of a simulation based on the [abm.AgentBase.Position].
type Spatial2D struct {
	plotcore.Plot

	// Sim is the simulation that this 2D representation is based on.
	Sim abm.Sim
}

func (sp *Spatial2D) Init() {
	sp.Plot.Init()
	sp.Styler(func(s *styles.Style) {
		s.Grow.Set(1, 1)
	})

	sp.Updater(func() {
		sp.SetPlot(sp.MakePlot()) // TODO: more optimized updating?
	})
}

// MakePlot creates the plot for the simulation.
func (sp *Spatial2D) MakePlot() *plot.Plot {
	pl := plot.New()

	agents := sp.Sim.Base().Agents

	xs := tensor.NewFloat32(len(agents))
	ys := tensor.NewFloat32(len(agents))

	for i, a := range agents {
		pos := a.Base().Position
		xs.Set(pos.X, i)
		ys.Set(pos.Y, i)
	}

	plot.Styler(ys, func(s *plot.Style) {
		s.Point.ColorFunc = func(i int) image.Image {
			beliefs := agents[i].Base().Beliefs
			var c hct.HCT
			switch len(beliefs) {
			case 1:
				c = hct.New(beliefs[0]*360, 100, 50)
			case 2:
				c = hct.New(beliefs[0]*360, beliefs[1]*100, 50)
			default:
				c = hct.New(beliefs[0]*360, beliefs[1]*100, beliefs[2]*100)
			}
			return colors.Uniform(c.AsRGBA())
		}
		s.Point.FillFunc = s.Point.ColorFunc
	})

	data := plot.Data{
		plot.X: xs,
		plot.Y: ys,
	}
	plots.NewScatter(pl, data)

	return pl
}
