package immutables

import (
	"path/filepath"

	"github.com/awe8128/arch-gen/utils/fs"
)

func OpenAPI(root string) error {
	path := filepath.Join(root, "openapi")
	filename := "config.yaml"
	content := `package: openapi
generate:
  gin-server: true
  strict-server: true
  models: true
  embedded-spec: true
output: ./be/presentation/openapi/openapi.gen.go
	`

	if err := fs.GenerateFile(content, path, filename); err != nil {
		return err
	}

	return nil
}
