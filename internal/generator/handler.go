package generator

import (
	"fmt"
	"path/filepath"

	"github.com/awe8128/arch-gen/templates"
	"github.com/awe8128/arch-gen/utils/fs"
)

func GenerateHandler(root string) error {

	path := filepath.Join(root, "presentation", "server")

	filename := "handler.go"

	content := fmt.Sprintf(
		`
		%s

		%s
		
		%s
		`,
		templates.Package("server"),
		templates.HandlerStruct(),
		templates.NewHandlerFunc(),
	)

	if err := fs.GenerateFile(content, path, filename); err != nil {
		return err
	}

	return nil
}
