// Copyright (c) 2025, Kleroterio. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package basic implements a basic agent-based model simulation
// for testing and demonstration purposes.
package basic

import "github.com/kleroterio/abm/abm"

// Sim is a basic agent-based model simulation.
type Sim struct {
	abm.SimBase
}

func (s *Sim) Init() {
	s.Agents = make([]abm.Agent, s.Config.(*Config).Population)
	for i := range s.Agents {
		s.Agents[i] = &abm.AgentBase{}
	}

	s.Base().Init()
}
