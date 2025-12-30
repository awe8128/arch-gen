package sys

import (
	"os"
	"path/filepath"

	"github.com/awe8128/arch-gen/config"
	"github.com/awe8128/arch-gen/internal/generator"

	"github.com/awe8128/arch-gen/utils/fs"
	"github.com/awe8128/arch-gen/utils/sqlc"
)

func Start() {

	// root folder
	root := config.GlobalConfig.Project.Name
	sys := config.GlobalConfig.Project.Sys
	// returns folder structure

	// create folder structure
	fs.Generate(sys)

	// infra layer
	for _, layer := range infrastructureLayer {
		path := filepath.Join(root, "infra", layer)
		if err := os.MkdirAll(path, 0o755); err != nil {
			panic(err)
		}

	}

	// db layer
	for _, layer := range DBLayer {
		path := filepath.Join(root, "infra", "db", layer)
		if err := os.MkdirAll(path, 0o755); err != nil {
			panic(err)
		}
	}

	// generate sqlc
	path := filepath.Join(root)

	content, filename := generator.SqlcYamlTemplate()
	if err := fs.GenerateFile(content, path, filename); err != nil {
		panic(err)
	}

	// domain layer
	for name, domain := range config.GlobalConfig.Domains {
		path := filepath.Join(root, "domain", name)

		if err := os.MkdirAll(path, 0o755); err != nil {
			panic(err)
		}

		content, filename := generator.EntityTemplate(name,
			domain.Properties,
		)

		if err := fs.GenerateFile(content, path, filename); err != nil {
			panic(err)
		}

		content, filename = generator.RepositoryTemplate(name,
			domain.Repositories,
		)

		if err := fs.GenerateFile(content, path, filename); err != nil {
			panic(err)
		}

		path = filepath.Join(root, "infra", "repository")
		content, filename = generator.InfraRepositoryTemplate(name,
			domain.Repositories,
			domain.Properties,
		)

		if err := fs.GenerateFile(content, path, filename); err != nil {
			panic(err)
		}

		path = filepath.Join(root, "infra", "db")
		content, filename = generator.StoreTemplate()
		if err := fs.GenerateFile(content, path, filename); err != nil {
			panic(err)
		}

	}
	tag := 1
	for table, columns := range config.GlobalConfig.DB {

		path = filepath.Join(root, "infra", "db", "migrations")
		content, filename = generator.MigrationTemplate(table, columns, tag)
		tag++

		if err := fs.GenerateFile(content, path, filename); err != nil {
			panic(err)
		}

		path = filepath.Join(root, "infra", "db", "query")
		content, filename = generator.QueryTemplate(table, columns)
		if err := fs.GenerateFile(content, path, filename); err != nil {
			panic(err)
		}

	}

	sqlc.RunSQLC()
}
