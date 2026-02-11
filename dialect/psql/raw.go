package psql

import (
	"github.com/aronsx/bob"
	"github.com/aronsx/bob/dialect/psql/dialect"
	"github.com/aronsx/bob/expr"
)

func RawQuery(q string, args ...any) bob.BaseQuery[expr.Clause] {
	return expr.RawQuery(dialect.Dialect, q, args...)
}
