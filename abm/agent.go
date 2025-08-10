// Copyright (c) 2025, Kleroterio. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package abm

// Agent is the interface that all agents implement.
type Agent interface {

	// AsBase returns the agent as an [AgentBase].
	AsBase() *AgentBase

	// Init initializes the agent with default values in the given simulation.
	Init(sim Sim)
}
