// Copyright (c) 2025, Kleroterio. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package abm

// ConfigBase is the base type for configuration parameter sets.
type ConfigBase struct {
}

func (c *ConfigBase) Base() *ConfigBase {
	return c
}
