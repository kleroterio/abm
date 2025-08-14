// Copyright (c) 2025, Kleroterio. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package abmcore

import (
	"math"

	"cogentcore.org/core/core"
	"cogentcore.org/core/styles"
	"cogentcore.org/core/tree"
	"cogentcore.org/lab/plotcore"
	"cogentcore.org/lab/stats/stats"
	"cogentcore.org/lab/table"
	"cogentcore.org/lab/tensor"
	"github.com/kleroterio/abm/abm"
)

// Stats is a customizable plot of statistics from a simulation.
type Stats struct {
	core.Frame

	// Sim is the simulation that this 2D representation is based on.
	Sim abm.Sim

	// statsTable is the stats data statsTable for plotting.
	statsTable *table.Table

	// agentTable is the table of agent data, which is used for computing statistics.
	agentTable *table.Table

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

func (st *Stats) makeStatsTable() {
	st.statsTable = table.New()
	st.statsTable.AddColumn("Polarization", tensor.NewFloat32(1))
}

// ComputeStats computes the statistics from the agent data.
func (st *Stats) ComputeStats() {
	if st.statsTable == nil {
		st.makeStatsTable()
	}

	steps := st.Sim.Base().Steps
	st.statsTable.SetNumRows(steps)

	varX := stats.VarPop(st.agentTable.Column("Belief X")).Float(0)
	varY := stats.VarPop(st.agentTable.Column("Belief Y")).Float(0)
	stddev := math.Sqrt(varX + varY)
	st.statsTable.Column("Polarization").SetFloat(stddev, steps-1)
}

func (st *Stats) MakeToolbar(p *tree.Plan) {
	st.plot.MakeToolbar(p)
}
