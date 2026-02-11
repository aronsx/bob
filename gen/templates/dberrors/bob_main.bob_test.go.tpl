{{$.Importer.Import "github.com/aronsx/bob"}}
// Set the testDB to enable tests that use the database
{{if eq $.Driver "github.com/jackc/pgx/v5" -}}
{{- $.Importer.Import "bobpgx" "github.com/aronsx/bob/drivers/pgx" -}}
var testDB bob.Transactor[bobpgx.Tx]
{{- else -}}
var testDB bob.Transactor[bob.Tx]
{{- end}}
