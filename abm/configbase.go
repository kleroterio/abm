// Copyright (c) 2025, Kleroterio. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package abm

// ConfigBase is the base type for configuration parameter sets.
type ConfigBase struct { //types:add

	// Beliefs is the number of political belief axes in the simulation.
	Beliefs int `default:"2"`

	// ChangeVelocity is the chance that an agent will change its spatial velocity.
	ChangeVelocity float32 `default:"0.1"`

	// VelocityMultiplier is a multiplier on the random spatial velocity at which
	// agents move.
	VelocityMultiplier float32 `default:"0.01"`

	// InteractionRadius is the multiplier on the maximum squared distance between
	// agents for an interaction to occur, with the base value being 1/n
	// (n = total number of agents).
	InteractionRadius float32 `default:"1"`

	// BeliefFilter is the impact that normalized belief distance has on the chance of
	// interaction. For example, a value of 1 means that if agents have a normalized
	// belief distance of 0.7, the chance of interaction is 30%. A value of 2 would
	// make that chance 15%. A value of 0 disables belief filtering.
	BeliefFilter float32 `default:"0.5"`

	// ExtremeBias indicates whether agents are biased towards extreme beliefs
	// (i.e., beliefs closer to 0 or 1 have a greater influence in interactions
	// than those closer to 0.5).
	ExtremeBias bool `default:"true"`

	// InteractionEffect is how much an interaction impacts beliefs as a
	// proportion of the initial difference in beliefs.
	InteractionEffect float32 `default:"0.01"`

	// ValueEffect is how much an agent's immutable values impact their beliefs
	// as a proportion of the difference between beliefs and values.
	// Values have a kind of restorative force, pulling beliefs back to the original
	// values over time.
	ValueEffect float32 `default:"0.005"`
}

func (cb *ConfigBase) Base() *ConfigBase {
	return cb
}

func (cb *ConfigBase) Defaults() {}
