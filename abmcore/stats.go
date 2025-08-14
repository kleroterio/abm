// Copyright (c) 2025, Kleroterio. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package abmcore

import (
	"cogentcore.org/core/core"
	"cogentcore.org/core/styles"
	"cogentcore.org/core/tree"
	"cogentcore.org/lab/plotcore"
	"cogentcore.org/lab/table"
	"github.com/kleroterio/abm/abm"
)

// Stats is a customizable plot of statistics from a simulation.
type Stats struct {
	core.Frame

	// Sim is the simulation that this 2D representation is based on.
	Sim abm.Sim

	// table is the data table for plotting.
	table *table.Table

	// plot is the plot editor widget.
	plot *plotcore.Editor
}

func (st *Stats) Init() {
	st.Frame.Init()
	st.Styler(func(s *styles.Style) {
		s.Grow.Set(1, 1)
		s.Direction = styles.Column
	})

	tree.AddChild(st, func(w *core.Toolbar) {
		w.Maker(st.MakeToolbar)
	})
	tree.AddChild(st, func(w *plotcore.Editor) {
		st.plot = w
	})
}

func (st *Stats) MakeToolbar(p *tree.Plan) {
	st.plot.MakeToolbar(p)
}
