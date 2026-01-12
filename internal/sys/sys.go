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

	// application layer
	for _, layer := range applicationLayer {
		path := filepath.Join(root, "application", layer)
		if err := os.MkdirAll(path, 0o755); err != nil {
			panic(err)
		}
	}

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

	path = filepath.Join(root, "config")
	content, filename = generator.ConfigTemplate()
	if err := fs.GenerateFile(content, path, filename); err != nil {
		panic(err)
	}

	// controller
	for _, folder := range PresentationLayer {
		path = filepath.Join(root, "presentation", folder)
		if err := os.MkdirAll(path, 0o755); err != nil {
			panic(err)
		}
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

		content, filename = generator.GenerateRepository(name,
			domain.Repositories,
		)

		if err := fs.GenerateFile(content, path, filename); err != nil {
			panic(err)
		}

		// infra repository
		path = filepath.Join(root, "infra", "repository")
		content, filename = generator.GenerateInfraRepository(name,
			domain.Repositories,
			domain.Properties,
		)

		if err := fs.GenerateFile(content, path, filename); err != nil {
			panic(err)
		}

		// infra db
		path = filepath.Join(root, "infra", "db")
		content, filename = generator.StoreTemplate()
		if err := fs.GenerateFile(content, path, filename); err != nil {
			panic(err)
		}

		// usecase
		path = filepath.Join(root, "application", "usecase", name)
		if err := os.MkdirAll(path, 0o755); err != nil {
			panic(err)
		}

		content, filename = generator.UsecaseTemplate(name, domain.Repositories)
		if err := fs.GenerateFile(content, path, filename); err != nil {
			panic(err)
		}

		// controller
		path = filepath.Join(root, "presentation", "controller")

		content, filename = generator.ControllerTemplate(name, domain.Repositories)
		if err := fs.GenerateFile(content, path, filename); err != nil {
			panic(err)
		}
		// di
		path = filepath.Join(root, "di")
		content, filename = generator.NewDI(name)
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

	path = filepath.Join(root, "infra", "db")
	content, filename = generator.InitDBTemplate()
	if err := fs.GenerateFile(content, path, filename); err != nil {
		panic(err)
	}

	path = filepath.Join(root, "presentation", "server")
	content, filename = generator.GenerateServer()
	if err := fs.GenerateFile(content, path, filename); err != nil {
		panic(err)
	}

	content, filename = generator.GenerateHandler()
	if err := fs.GenerateFile(content, path, filename); err != nil {
		panic(err)
	}
}
