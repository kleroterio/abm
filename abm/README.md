# abm

This document explains the architecture of **abm**.

## Agent

All agents embed `AgentBase`. An agent can make a variety of decisions. For each decision, it evaluates the utility of each option, and then picks the option with the highest utility.

Utility is computed as a function of three factors, personal, social, and moral, the importance of which is controlled by parameters of the same name.

* Personal factors consist of things like money and power.
* Social factors consist of the weighted average utility for social connections.
* Moral factors consist of the perceived median utility and inequality.
