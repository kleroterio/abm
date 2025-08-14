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

// Step advances the simulation by one time step.
// It does this by having each agent interact with one or more randomly
// selected agents as determined by the configuration parameters.
func (sb *SimBase) Step() {
	cb := sb.Config.Base()
	ir := cb.InteractionRadius / float32(len(sb.Agents))
	for i, a := range sb.Agents {
		a.Base().StepPosition()
		a.Base().ApplyValues()
		for j, other := range sb.Agents {
			if i == j {
				continue
			}
			dist := a.Base().Position.DistanceToSquared(other.Base().Position)
			if dist > ir {
				continue
			}

			if cb.BeliefFilter > 0 {
				beliefDist := float32(0)
				for i, ba := range a.Base().Beliefs {
					delta := other.Base().Beliefs[i] - ba
					beliefDist += delta * delta
				}
				beliefDist = math32.Sqrt(beliefDist / float32(cb.Beliefs))
				chanceInteract := (1 - beliefDist) / cb.BeliefFilter
				if rand.Float32() > chanceInteract {
					continue
				}
			}

			a.Base().Interact(other)
			break
		}
	}
}
