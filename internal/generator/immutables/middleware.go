package immutables

import (
	"path/filepath"

	"github.com/awe8128/arch-gen/templates"
	"github.com/awe8128/arch-gen/utils/fs"
)

func GenerateMiddleware(root string) error {

	path := filepath.Join(root, "presentation", "server", "middleware")

	filename, content := templates.MiddlewareTemplate()

	if err := fs.GenerateFile(content, path, filename); err != nil {
		return err
	}

	return nil
}
