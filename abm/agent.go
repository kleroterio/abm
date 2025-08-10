// Copyright (c) 2025, Kleroterio. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package abm

// Agent is the interface that all agents implement.
type Agent interface {

	// AsAgentBase returns the agent as an [AgentBase].
	AsAgentBase() *AgentBase

	// Init initializes the agent with default values.
	Init()
}
