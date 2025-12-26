package structure

import (
	"github.com/awe8128/arch-gen/config"
	"github.com/awe8128/arch-gen/internal/files"
	"github.com/awe8128/arch-gen/internal/helper"
)

func Start() {
	// domain layer
	for name, domain := range config.GlobalConfig.Domains {

		content, dir, filename := helper.GenerateEntityContent(name,
			domain.Properties,
		)

		if err := files.Generate(content, dir, filename); err != nil {
			panic(err)
		}

		content, dir, filename = helper.GenerateRepositoryContent(name,
			domain.Repositories,
		)

		if err := files.Generate(content, dir, filename); err != nil {
			panic(err)
		}
	}
}
