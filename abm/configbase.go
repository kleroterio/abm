// Copyright (c) 2025, Kleroterio. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package abm

// ConfigBase is the base type for configuration parameter sets.
type ConfigBase struct { //types:add

	// Beliefs is the number of belief axes in the simulation.
	Beliefs int `default:"2"`
}

func (cb *ConfigBase) Base() *ConfigBase {
	return cb
}

func (cb *ConfigBase) Defaults() {}
