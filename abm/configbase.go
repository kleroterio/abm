// Copyright (c) 2025, Kleroterio. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package abm

// ConfigBase is the base type for configuration parameter sets.
type ConfigBase struct { //types:add

	// Beliefs is the number of political belief axes in the simulation.
	Beliefs int `default:"2"`

	// SpatialNeighbors is the number of agents each agent will be connected to
	// by virtue of their spatial proximity.
	SpatialNeighbors int `default:"5"`

	// BeliefNeighbors is the number of agents each agent will be connected to
	// by virtue of their proximity in political beliefs.
	BeliefNeighbors int `default:"5"`
}

func (cb *ConfigBase) Base() *ConfigBase {
	return cb
}

func (cb *ConfigBase) Defaults() {}
