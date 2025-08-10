// Copyright (c) 2025, Kleroterio. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package abm

// Sim is the interface that all simulations implement.
type Sim interface {

	// Base returns the simulation as a [SimBase].
	Base() *SimBase
}
