help: ## Show this help
	@grep -E '^[a-zA-Z0-9_.-]+:.*?## ' Makefile | sort \
	  | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[1m%-15s\033[0m %s\n", $$1, $$2}'
	  
run:
	go run main.go -c arch.yaml
	
server:
	cd be && go run cmd/api/main.go

oapi: ## generate oapi code
	go tool oapi-codegen --config ./openapi/config.yaml ./openapi/openapi.gen.yaml

.PHONY: all