// Copyright (c) 2025, Kleroterio. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package basic

import "github.com/kleroterio/abm/abm"

type Config struct {

	// Population is the number of citizens in the simulation.
	Population int `default:"100"`

	abm.ConfigBase
}
