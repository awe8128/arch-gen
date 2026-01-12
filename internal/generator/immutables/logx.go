package immutables

import (
	"path/filepath"

	"github.com/awe8128/arch-gen/templates"
	"github.com/awe8128/arch-gen/utils/fs"
)

func GenerateLogx(root string) error {
	path := filepath.Join(root, "presentation", "server", "logx")

	filename, content := templates.LogxTemplate()
	if err := fs.GenerateFile(content, path, filename); err != nil {
		return err
	}

	filename, content = templates.CustomCodeTemplate()
	if err := fs.GenerateFile(content, path, filename); err != nil {
		return err
	}

	filename, content = templates.HelperTemplate()
	if err := fs.GenerateFile(content, path, filename); err != nil {
		return err
	}

	return nil
}
