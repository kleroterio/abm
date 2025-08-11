// Copyright (c) 2025, Kleroterio. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package abm

// Config is the interface that all configuration parameter sets implement.
type Config interface {

	// Base returns the configuration as a [ConfigBase].
	Base() *ConfigBase
}
