package generator

import (
	"path/filepath"

	"github.com/awe8128/arch-gen/utils/fs"
)

func GenerateNewAPI(root string) error {
	path := filepath.Join(root, "cmd", "api")

	filename := "main.go"
	content := `
	package main
	func main() {
		ctx := context.Background()
		config, err := config.Load(".")
		if err != nil {
			log.Printf("Failed to load config: %v", err)
		}

		store, err := db.Init(ctx, config)
		if err != nil {
			log.Fatalf("Failed to load: %v", err)
		}

		handler := server.NewHandler(store)

		server, err := server.NewServer(config, handler)
		if err != nil {
			log.Fatalf("Failed to init server: %v", err)
		}

		server.Run(config.API_PORT)
	}
	`
	if err := fs.GenerateFile(content, path, filename); err != nil {
		return err
	}
	return nil
}
