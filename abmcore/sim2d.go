// Copyright (c) 2025, Kleroterio. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package abmcore

import (
	"cogentcore.org/core/core"
	"cogentcore.org/core/tree"
	"github.com/kleroterio/abm/abm"
)

// Sim2D implements a plot-based 2D representation of an agent-based model simulation.
type Sim2D struct {
	core.Frame

	// Sim is the simulation that this 2D representation is based on.
	Sim abm.Sim
}

func (sw *Sim2D) Init() {
	sw.Frame.Init()

	tree.AddChild(sw, func(w *core.Tabs) {
		w.NewTab("Spatial")
	})
}
