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

	n := len(agents)
	spatialX := tensor.NewFloat32(n)
	spatialY := tensor.NewFloat32(n)
	beliefX := tensor.NewFloat32(n)
	beliefY := tensor.NewFloat32(n)

	for i, a := range agents {
		pos := a.Base().Position
		spatialX.Set(pos.X, i)
		spatialY.Set(pos.Y, i)

		beliefs := a.Base().Beliefs
		beliefX.Set(beliefs[0], i)
		beliefY.Set(beliefs[1], i)
	}

	plot.Styler(spatialY, func(s *plot.Style) {
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
		plot.X: spatialX,
		plot.Y: spatialY,
	}
	plots.NewScatter(pl, data)

	return pl
}
