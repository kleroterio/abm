// Copyright (c) 2025, Kleroterio. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package abm

// Sim is the interface that all simulations implement.
type Sim interface {

	// Base returns the simulation as a [SimBase].
	Base() *SimBase

	// Init initializes the simulation with default values.
	Init()
}

// NewSim creates and initializes a new simulation of type S.
// *S must implement the [Sim] interface.
func NewSim[S any]() *S {
	simS := new(S)
	sim := any(simS).(Sim)
	sim.Base().This = sim
	sim.Init()
	return simS
}
