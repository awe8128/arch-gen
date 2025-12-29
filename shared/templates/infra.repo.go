package templates

import "fmt"

func NewInfraTemplate(name string) string {
	content := fmt.Sprintf(
		`
type %sStore struct {
	db *db.SQLStore
	attrs []slog.Attr
}

func New%sRepository(db *db.SQLStore) %sRepository {
	return &%sStore{
		db: db,
		attrs: []slog.Attr{
			slog.String("layer", "repository"),
			slog.String("domain", "task"),
		},
	}
}
`,
		capitalize(name), capitalize(name), capitalize(name), capitalize(name),
	)

	return content

}
