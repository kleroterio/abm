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
		NewPlot(fr).SetSim(sw.Sim)
	})
}

func (sw *Sim2D) MakeToolbar(p *tree.Plan) {
	tree.Add(p, func(w *core.Button) {
		w.SetText("Reset").SetIcon(icons.Update)
		w.OnClick(func(e events.Event) {
			sw.Sim.Init()
			sw.Update()
		})
	})
}
