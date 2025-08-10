// Copyright (c) 2025, Kleroterio. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package abm

import "cogentcore.org/core/math32"

// AgentBase is the base type for all agents.
type AgentBase struct {

	// ID is the unique identifier for the agent.
	ID uint

	// Position is the current position of the agent in the simulation space.
	Position math32.Vector2

	// Connections holds the connections between this agent and others.
	// The key is the ID of the connected agent, and the value is the strength of
	// the connection (-1 to 1), with negative values indicating an oppositional
	// connection.
	Connections map[uint]float32

	// Beliefs contains the agent's beliefs on each belief axis (0 to 1).
	Beliefs []float32
}
