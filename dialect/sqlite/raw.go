package sqlite

import (
	"github.com/aronsx/bob"
	"github.com/aronsx/bob/dialect/sqlite/dialect"
	"github.com/aronsx/bob/expr"
)

func RawQuery(q string, args ...any) bob.BaseQuery[expr.Clause] {
	return expr.RawQuery(dialect.Dialect, q, args...)
}
