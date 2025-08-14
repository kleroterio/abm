// Copyright (c) 2025, Kleroterio. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package abm

import (
	"math/rand/v2"
	"slices"
	"sync/atomic"

	"cogentcore.org/core/math32"
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

	// Velocity is the current velocity of the agent in the simulation space.
	Velocity math32.Vector2

	// Connections holds the connections between this agent and others.
	// The key is the ID of the connected agent, and the value is the strength of
	// the connection (-1 to 1), with negative values indicating an oppositional
	// connection.
	// Connections map[uint]float32

	// Beliefs contains the agent's beliefs on each belief axis (0 to 1).
	Beliefs []float32

	// Values contains the agent's values on each belief axis (0 to 1).
	// Each value corresponds to the belief on the same axis, but values
	// are immutable unlike beliefs.
	Values []float32

	// Influence is the agent's influence on others in the simulation
	// (initial value 0 to 1).
	Influence float32
}

func (ab *AgentBase) Base() *AgentBase {
	return ab
}

// Init initializes the agent with default values.
func (ab *AgentBase) Init(sim Sim) {
	sb := sim.Base()
	cb := sb.Config.Base()

	ab.Sim = sim
	ab.ID = atomic.AddUint64(&sb.idCounter, 1) - 1

	ab.Beliefs = make([]float32, cb.Beliefs)
	for i := range ab.Beliefs {
		ab.Beliefs[i] = rand.Float32()
	}
	ab.Values = slices.Clone(ab.Beliefs)
	ab.Influence = cb.RandomInfluence*rand.Float32() + (1 - cb.RandomInfluence)
	if cb.PartisanPosition && cb.Beliefs >= 2 {
		ab.Position.Set(ab.Beliefs[0], ab.Beliefs[1])
	} else {
		ab.Position = math32.Vec2(rand.Float32(), rand.Float32())
	}
}

var zeroVec, oneVec = math32.Vector2{}, math32.Vec2(1, 1)

// StepPosition updates the agent's position and velocity one time step.
func (ab *AgentBase) StepPosition() {
	cb := ab.Sim.Base().Config.Base()
	if rand.Float32() < cb.ChangeVelocity {
		ab.Velocity = math32.Vec2(rand.Float32(), rand.Float32()).SubScalar(0.5).MulScalar(1 - cb.BeliefVelocity)
		if cb.Beliefs >= 2 {
			ab.Velocity.SetAdd(math32.Vec2(ab.Beliefs[0], ab.Beliefs[1]).SubScalar(0.5).MulScalar(cb.BeliefVelocity))
		}
		ab.Velocity.SetMulScalar(cb.VelocityMultiplier)
	}
	ab.Position.SetAdd(ab.Velocity)
	ab.Position.Clamp(zeroVec, oneVec)
}

// ApplyValues applies the restorative effect of the agent's values on its beliefs.
func (ab *AgentBase) ApplyValues() {
	cb := ab.Sim.Base().Config.Base()
	for i := range ab.Beliefs {
		ba := &ab.Beliefs[i]
		*ba += cb.ValueEffect * (ab.Values[i] - *ba)
		*ba = math32.Clamp(*ba, 0, 1)
	}
}

// Interact has the agent interact with the given other agent.
func (ab *AgentBase) Interact(other Agent) {
	cb := ab.Sim.Base().Config.Base()
	ai := ab.Influence
	oi := other.Base().Influence
	for i := range ab.Beliefs {
		ba := &ab.Beliefs[i]
		bo := &other.Base().Beliefs[i]

		ab.shiftBelief(ba, bo, ai, oi, cb)
		other.Base().shiftBelief(bo, ba, oi, ai, cb) // reverse interaction

		*ba = math32.Clamp(*ba, 0, 1)
		*bo = math32.Clamp(*bo, 0, 1)
	}
}

// shiftBelief shifts the belief ba towards bo based on the influences ai and oi.
func (ab *AgentBase) shiftBelief(ba, bo *float32, ai, oi float32, cb *ConfigBase) {
	// The delta is not just based on bo - ba, because talking with someone who
	// you agree with will move you more strongly in that direction, so it is more
	// like bo - 0.5. On the other hand, arguments in the middle aren't entirely
	// un-motivating, just somewhat less persuasive, so the ExtremeBias parameter
	// determines how much less persuasive they are.
	baseline := 0.5*cb.ExtremeBias + *ba*(1-cb.ExtremeBias)
	delta := cb.InteractionEffect * (*bo - baseline)
	*ba += delta * (oi / ai)
}
