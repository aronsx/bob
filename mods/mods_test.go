package mods

import (
	"github.com/aronsx/bob"
	"github.com/aronsx/bob/clause"
)

var (
	_ bob.Mod[any]                                = QueryMods[any](nil)
	_ bob.Mod[interface{ AppendCTE(clause.CTE) }] = With[interface{ AppendCTE(clause.CTE) }]{}
)
