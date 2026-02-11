{{- define "helpers/where_variables"}}
{{$.Importer.Import (printf "github.com/aronsx/bob/dialect/%s/dialect" $.Dialect)}}
var (
	SelectWhere = Where[*dialect.SelectQuery]()
	UpdateWhere = Where[*dialect.UpdateQuery]()
	DeleteWhere = Where[*dialect.DeleteQuery]()
)
{{- end -}}
