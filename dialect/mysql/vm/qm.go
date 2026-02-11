package vm

import (
	"github.com/aronsx/bob"
	"github.com/aronsx/bob/clause"
	"github.com/aronsx/bob/dialect/mysql/dialect"
	"github.com/aronsx/bob/mods"
)

func RowValue(clauses ...bob.Expression) bob.Mod[*dialect.ValuesQuery] {
	return mods.Values[*dialect.ValuesQuery](clauses)
}

func OrderBy(e any) dialect.OrderBy[*dialect.ValuesQuery] {
	return dialect.OrderBy[*dialect.ValuesQuery](func() clause.OrderDef {
		return clause.OrderDef{
			Expression: e,
		}
	})
}

func Limit(count int64) bob.Mod[*dialect.ValuesQuery] {
	return mods.Limit[*dialect.ValuesQuery]{
		Count: count,
	}
}
