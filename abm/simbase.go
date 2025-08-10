// Copyright (c) 2025, Kleroterio. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package abm

// SimBase is the base type for all simulations.
type SimBase struct {

	// Agents are the agents in the simulation.
	Agents []Agent

	// idCounter is used to generate unique IDs for agents.
	idCounter uint64
}

func (sb *SimBase) AsBase() *SimBase {
	return sb
}
