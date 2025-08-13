// Copyright (c) 2025, Kleroterio. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package abm

import (
	"math/rand/v2"
	"sync/atomic"

	"cogentcore.org/core/math32"
	"cogentcore.org/lab/tensor"
)

// AgentBase is the base type for all agents.
type AgentBase struct {

	// Sim is the simulation that the agent belongs to.
	Sim Sim

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

func (ab *AgentBase) Base() *AgentBase {
	return ab
}

// Init initializes the agent with default values.
func (ab *AgentBase) Init(sim Sim) {
	sb := sim.Base()

	ab.Sim = sim
	ab.ID = atomic.AddUint64(&sb.idCounter, 1) - 1

	ab.Position = math32.Vec2(rand.Float32(), rand.Float32())
	ab.Beliefs = make([]float32, sb.Config.Base().Beliefs)
	for i := range ab.Beliefs {
		ab.Beliefs[i] = rand.Float32()
	}
}

// Tensor returns the agent's beliefs and position as a tensor.
func (ab *AgentBase) Tensor() *tensor.Float32 {
	return tensor.NewFloat32FromValues(append(ab.Beliefs, ab.Position.X, ab.Position.Y)...)
}

// Interact has the agent interact with the given other agent.
func (ab *AgentBase) Interact(other Agent) {
	ie := ab.Sim.Base().Config.Base().InteractionEffect
	for i, b := range ab.Beliefs {
		delta := ie * (other.Base().Beliefs[i] - b)
		ab.Beliefs[i] += delta
		other.Base().Beliefs[i] -= delta
	}
}
