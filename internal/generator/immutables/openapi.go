package immutables

import "path/filepath"

func OpenAPI(root string) (string, string, string) {
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
	return path, filename, content
}
