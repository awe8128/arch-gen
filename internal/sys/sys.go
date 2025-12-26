package sys

import (
	"os"
	"path/filepath"

	"github.com/awe8128/arch-gen/config"
	"github.com/awe8128/arch-gen/internal/generator"

	"github.com/awe8128/arch-gen/utils/fs"
)

func Start() {

	// root folder
	root := config.GlobalConfig.Project.Name
	sys := config.GlobalConfig.Project.Sys
	// returns folder structure

	// create folder structure
	fs.Generate(sys)

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
	}
}
