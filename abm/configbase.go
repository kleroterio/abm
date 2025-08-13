// Copyright (c) 2025, Kleroterio. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package abm

// ConfigBase is the base type for configuration parameter sets.
type ConfigBase struct { //types:add

	// Beliefs is the number of political belief axes in the simulation.
	Beliefs int `default:"2"`

	// Interactions is the number of interactions per agent per step.
	Interactions int `default:"2"`

	// InteractionEffect is how much an interaction impacts beliefs as a
	// proportion of the initial difference in beliefs.
	InteractionEffect float32 `default:"0.01"`
}

func (cb *ConfigBase) Base() *ConfigBase {
	return cb
}

func (cb *ConfigBase) Defaults() {}
