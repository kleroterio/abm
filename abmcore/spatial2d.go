// Copyright (c) 2025, Kleroterio. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package abmcore

import (
	"cogentcore.org/lab/plotcore"
	"github.com/kleroterio/abm/abm"
)

// Spatial2D is a 2d plot of a simulation based on the [abm.AgentBase.Position].
type Spatial2D struct {
	plotcore.Editor

	// Sim is the simulation that this 2D representation is based on.
	Sim abm.Sim
}

func (sp *Spatial2D) Init() {
	sp.Editor.Init()
}
