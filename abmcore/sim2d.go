// Copyright (c) 2025, Kleroterio. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package abmcore

import (
	"cogentcore.org/core/core"
	"cogentcore.org/core/events"
	"cogentcore.org/core/icons"
	"cogentcore.org/core/styles"
	"cogentcore.org/core/tree"
	"github.com/kleroterio/abm/abm"
)

// Sim2D implements a plot-based 2D representation of an agent-based model simulation.
type Sim2D struct {
	core.Splits

	// Sim is the simulation that this 2D representation is based on.
	Sim abm.Sim

	// running is whether the simulation is currently running.
	running bool

	// population is the plot of the agent population.
	population *Plot
}

func (sw *Sim2D) Init() {
	sw.Splits.Init()
	sw.SetSplits(0.25, 0.75)
	sw.Styler(func(s *styles.Style) {
		s.Grow.Set(1, 1)
	})

	tree.AddChild(sw, func(w *core.Form) {
		w.SetStruct(sw.Sim.Base().Config)
	})
	tree.AddChild(sw, func(w *core.Tabs) {
		fr, _ := w.NewTab("Population 2D")
		sw.population = NewPlot(fr).SetSim(sw.Sim)
	})
}

func (sw *Sim2D) MakeToolbar(p *tree.Plan) {
	tree.Add(p, func(w *core.Button) {
		w.SetText("Reset").SetIcon(icons.Update)
		w.OnClick(func(e events.Event) {
			sw.running = false
			sw.Sim.Init()
			sw.population.UpdatePlot()
			sw.Scene.Restyle()
		})
	})
	tree.Add(p, func(w *core.Button) {
		w.SetText("Run").SetIcon(icons.PlayArrow)
		w.OnClick(func(e events.Event) {
			sw.running = true
			sw.Scene.Restyle()
			sw.Animate(func(a *core.Animation) {
				if !sw.running {
					a.Done = true
					return
				}
				sw.Sim.Step()
				sw.population.UpdatePlot()
			})
		})
	})
	tree.Add(p, func(w *core.Button) {
		w.SetText("Stop").SetIcon(icons.Stop)
		w.OnClick(func(e events.Event) {
			sw.running = false
			w.Restyle()
		})
		w.FirstStyler(func(s *styles.Style) {
			s.SetEnabled(sw.running)
		})
	})
	tree.Add(p, func(w *core.Button) {
		w.SetText("Step").SetIcon(icons.Step)
		w.OnClick(func(e events.Event) {
			sw.Sim.Step()
			sw.population.UpdatePlot()
		})
	})
}
