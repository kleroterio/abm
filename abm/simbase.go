// Copyright (c) 2025, Kleroterio. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package abm

// SimBase is the base type for all simulations.
type SimBase struct {

	// This is the pointer to the simulation in the actual non-base
	// type of the simulation.
	This Sim

	// Agents are the agents in the simulation.
	Agents []Agent

	// idCounter is used to generate unique IDs for agents.
	idCounter uint64

	// NumBeliefs is the number of belief axes in the simulation.
	// This defaults to 2.
	NumBeliefs int
}

func (sb *SimBase) Base() *SimBase {
	return sb
}

// Init initializes the simulation by initializing all agents
// and connecting them according to their positions and beliefs.
func (sb *SimBase) Init() {
	if sb.NumBeliefs <= 0 {
		sb.NumBeliefs = 2
	}

	for _, a := range sb.Agents {
		a.Init(sb.This)
	}
}
