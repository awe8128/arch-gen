package immutables

func SqlcYamlTemplate() (string, string) {
	filename := "sqlc.yaml"

	content :=
		`version: "2"
sql:
- engine: "postgresql"
queries: "./infra/db/query/"
schema: "./infra/db/migrations/"
gen:
  go:
	package: "sqlc"
	out: "./infra/db/sqlc"
	sql_package: "pgx/v5"
	emit_json_tags: true
	emit_prepared_queries: false
	emit_interface: true
	emit_exact_table_names: true
	emit_empty_slices: true
	overrides:
	  - db_type: "timestamptz"
		go_type: "time.Time"
	  - db_type: "uuid"
		go_type: "github.com/google/uuid.UUID"
`
	return content, filename
}
