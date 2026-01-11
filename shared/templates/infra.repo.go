package templates

import "fmt"

func StoreTemplate(name string) string {
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
			slog.String("domain", "%s"),
		},
	}
}
`,
		capitalize(name), capitalize(name), capitalize(name), capitalize(name), name,
	)

	return content

}
