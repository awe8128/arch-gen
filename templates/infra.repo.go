package templates

import (
	"fmt"

	"github.com/awe8128/arch-gen/templates/utils"
)

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
		utils.Capitalize(name), utils.Capitalize(name), utils.Capitalize(name), utils.Capitalize(name), name,
	)

	return content

}
