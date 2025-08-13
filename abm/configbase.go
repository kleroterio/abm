// Copyright (c) 2025, Kleroterio. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package abm

// ConfigBase is the base type for configuration parameter sets.
type ConfigBase struct { //types:add

	// Beliefs is the number of political belief axes in the simulation.
	Beliefs int `default:"2"`

	// ChangeVelocity is the chance that an agent will change its velocity.
	ChangeVelocity float32 `default:"0.1"`

	// SpatialSpeed is the speed at which agents randomly move spatially.
	SpatialSpeed float32 `default:"0.01"`

	// Interactions is the number of interactions per agent per step.
	Interactions int `default:"2"`

	// SpatialWeight is the importance of spatial proximity in determing who to
	// interact with.
	SpatialWeight float32 `default:"1"`

	// BeliefWeight is the importance of beliefs in determing who to interact with.
	BeliefWeight float32 `default:"1"`

	// InteractionRadius is the maximum distance between agents for an
	// interaction to occur.
	InteractionRadius float32 `default:"0.4"`

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
