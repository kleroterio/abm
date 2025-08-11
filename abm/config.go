// Copyright (c) 2025, Kleroterio. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package abm

import (
	"cogentcore.org/core/base/errors"
	"cogentcore.org/core/base/reflectx"
)

// Config is the interface that all configuration parameter sets implement.
type Config interface {

	// Base returns the configuration as a [ConfigBase].
	Base() *ConfigBase

	// Defaults sets any special default values for the configuration.
	// Defaults specified via `default:"..."` struct tags are set automatically
	// (in [NewConfig]).
	Defaults()
}

// NewConfig creates and initializes a new configuration of type C.
// *C must implement the [Config] interface.
func NewConfig[C any]() *C {
	cfgC := new(C)
	cfg := any(cfgC).(Config)
	errors.Log(reflectx.SetFromDefaultTags(cfg))
	cfg.Defaults()
	return cfgC
}
