// Copyright (c) 2025, Kleroterio. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package abm

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
