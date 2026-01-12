package immutables

import "path/filepath"

func Makefile(root string) (string, string, string) {

	path := filepath.Join(root)
	filename := "Makefile"
	content := `
help: ## Show this help
	@grep -E '^[a-zA-Z0-9_.-]+:.*?## ' Makefile | sort \
	| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[1m%-15s\033[0m %s\n", $$1, $$2}'

oapi: ## generate oapi code
	go tool oapi-codegen --config ./openapi/config.yaml ./openapi/openapi.gen.yaml

bundle: ## bundle openapi spec
	docker run --rm -v $$PWD:/spec redocly/cli bundle ./openapi/openapi.yaml -o openapi/openapi.gen.yaml

sqlc: ## generate sqlc code
	sqlc generate

.PHONY: all
`

	return path, filename, content
}
