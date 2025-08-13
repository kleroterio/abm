// Copyright (c) 2025, Kleroterio. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package abm

import (
	"cogentcore.org/core/base/errors"
	"github.com/mroth/weightedrand/v2"
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
	for i, a := range sb.Agents {
		at := a.Base().Tensor()
		choices := []weightedrand.Choice[int, int]{}
		for j, other := range sb.Agents {
			if i == j {
				continue
			}
			weight := a.Base().InteractionWeight(at, other)
			if weight == 0 {
				continue
			}
			choices = append(choices, weightedrand.NewChoice(j, weight))
		}
		if len(choices) == 0 {
			continue // no one to interact with
		}
		chooser, err := weightedrand.NewChooser(choices...)
		if errors.Log(err) != nil {
			continue
		}

		for range sb.Config.Base().Interactions {
			j := chooser.Pick()
			other := sb.Agents[j]
			a.Base().Interact(other)
		}
	}
}
