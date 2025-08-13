// Copyright (c) 2025, Kleroterio. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package abm

import (
	"math/rand/v2"

	"cogentcore.org/core/math32"
)

// SimBase is the base type for all simulations.
type SimBase struct {

	// This is the pointer to the simulation in the actual non-base
	// type of the simulation.
	This Sim

	// Config has the configuration parameters for the simulation.
	Config Config

	// Agents are the agents in the simulation.
	Agents []Agent

	// idCounter is used to generate unique IDs for agents.
	idCounter uint64
}

func (sb *SimBase) Base() *SimBase {
	return sb
}

// Init initializes the simulation by initializing all agents
// and connecting them according to their positions and beliefs.
func (sb *SimBase) Init() {
	for _, a := range sb.Agents {
		a.Init(sb.This)
	}
}

var zeroVec, oneVec = math32.Vector2{}, math32.Vec2(1, 1)

// Step advances the simulation by one time step.
// It does this by having each agent interact with one or more randomly
// selected agents as determined by the configuration parameters.
func (sb *SimBase) Step() {
	for i, a := range sb.Agents {
		delta := math32.Vec2(rand.Float32(), rand.Float32()).SubScalar(0.5).MulScalar(sb.Config.Base().SpatialSpeed)
		a.Base().Position.SetAdd(delta)
		a.Base().Position.Clamp(zeroVec, oneVec)
		for j, other := range sb.Agents {
			if i == j {
				continue
			}
			dist := a.Base().Position.DistanceToSquared(other.Base().Position)
			if dist < 0.01 {
				a.Base().Interact(other)
				break
			}
		}
	}
}
