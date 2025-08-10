// Copyright (c) 2025, Kleroterio. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package abm

import (
	"math/rand/v2"
	"sync/atomic"

	"cogentcore.org/core/math32"
)

// idCounter is used to generate unique IDs for agents.
var idCounter uint64

// AgentBase is the base type for all agents.
type AgentBase struct {

	// ID is the unique identifier for the agent.
	ID uint64

	// Position is the current normalized position of the agent in the simulation space
	// (each axis is from 0 to 1).
	Position math32.Vector2

	// Connections holds the connections between this agent and others.
	// The key is the ID of the connected agent, and the value is the strength of
	// the connection (-1 to 1), with negative values indicating an oppositional
	// connection.
	Connections map[uint]float32

	// Beliefs contains the agent's beliefs on each belief axis (0 to 1).
	Beliefs []float32
}

func (ab *AgentBase) AsAgentBase() *AgentBase {
	return ab
}

// Init initializes the agent with default values.
func (ab *AgentBase) Init() {
	ab.ID = atomic.AddUint64(&idCounter, 1) - 1
	ab.Position = math32.Vec2(rand.Float32(), rand.Float32())
}
