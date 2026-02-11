package vm

import (
	"github.com/aronsx/bob"
	"github.com/aronsx/bob/dialect/sqlite/dialect"
	"github.com/aronsx/bob/mods"
)

func RowValue(clauses ...bob.Expression) bob.Mod[*dialect.ValuesQuery] {
	return mods.Values[*dialect.ValuesQuery](clauses)
}
