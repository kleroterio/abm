// Copyright (c) 2025, Kleroterio. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"cogentcore.org/core/core"
	"github.com/kleroterio/abm/abm"
	"github.com/kleroterio/abm/abmcore"
	"github.com/kleroterio/abm/sims/basic"
)

func main() {
	b := core.NewBody("Basic ABM Simulation")
	sim := abm.NewSim[basic.Sim]()
	sw := abmcore.NewSim2D(b).SetSim(sim)
	b.AddTopBar(func(bar *core.Frame) {
		core.NewToolbar(bar).Maker(sw.MakeToolbar)
	})
	b.RunMainWindow()
}
