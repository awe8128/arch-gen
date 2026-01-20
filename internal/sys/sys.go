package sys

import (
	"os"
	"path/filepath"

	"github.com/awe8128/arch-gen/config"
	"github.com/awe8128/arch-gen/internal/generator"
	"github.com/awe8128/arch-gen/internal/generator/immutables"

	"github.com/awe8128/arch-gen/utils/fs"
	"github.com/awe8128/arch-gen/utils/sqlc"
)

func Start() {

	// root folder
	root := config.GlobalConfig.Project.Name
	sys := config.GlobalConfig.Project.Sys

	// create folder structure
	fs.Generate(sys)

	// sqlc yaml config
	err := immutables.SqlcYamlTemplate(root)
	if err != nil {
		panic(err)
	}

	// config file
	err = immutables.ConfigTemplate(root)
	if err != nil {
		panic(err)
	}

	// domain layer
	for name, domain := range config.GlobalConfig.Domains {
		path := filepath.Join(root, "domain", name)
		if err := os.MkdirAll(path, 0o755); err != nil {
			panic(err)
		}

		content, filename := generator.GenerateEntity(name,
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
		content, filename = immutables.StoreTemplate()
		if err := fs.GenerateFile(content, path, filename); err != nil {
			panic(err)
		}

		// usecase
		path = filepath.Join(root, "application", "usecase", name)
		if err := os.MkdirAll(path, 0o755); err != nil {
			panic(err)
		}

		content, filename = generator.GenerateUsecase(name, domain.Repositories)
		if err := fs.GenerateFile(content, path, filename); err != nil {
			panic(err)
		}

		// controller
		path = filepath.Join(root, "presentation", "controller")

		content, filename = generator.GenerateController(name, domain.Repositories)
		if err := fs.GenerateFile(content, path, filename); err != nil {
			panic(err)
		}
		// di
		path = filepath.Join(root, "di")
		content, filename = generator.GenerateDI(name)
		if err := fs.GenerateFile(content, path, filename); err != nil {
			panic(err)
		}
	}

	tag := 1
	for table, columns := range config.GlobalConfig.DB {

		path := filepath.Join(root, "infra", "db", "migrations")
		content, filename := generator.GenerateMigration(table, columns, tag)
		tag++

		if err := fs.GenerateFile(content, path, filename); err != nil {
			panic(err)
		}

		path = filepath.Join(root, "infra", "db", "query")
		content, filename = generator.GenerateQuery(table, columns)
		if err := fs.GenerateFile(content, path, filename); err != nil {
			panic(err)
		}

	}

	err = sqlc.RunSQLC()
	if err != nil {
		panic(err)
	}

	err = immutables.GenerateDBInit(root)
	if err != nil {
		panic(err)
	}

	err = immutables.GenerateServer(root)
	if err != nil {
		panic(err)
	}

	err = generator.GenerateHandler(root)
	if err != nil {
		panic(err)
	}

	err = generator.GenerateRouter(root)
	if err != nil {
		panic(err)
	}

	err = generator.GenerateNewAPI(root)
	if err != nil {
		panic(err)
	}

	err = immutables.OpenAPI(root)
	if err != nil {
		panic(err)
	}

	err = immutables.Makefile(root)
	if err != nil {
		panic(err)
	}

	err = immutables.GenerateLogx(root)
	if err != nil {
		panic(err)
	}

	err = immutables.GenerateMiddleware(root)
	if err != nil {
		panic(err)
	}

}
