package fs

import (
	"os"
	"path/filepath"

	"github.com/awe8128/arch-gen/config"
)

func Generate(systemDesign string) {
	switch systemDesign {
	case "ddd":
		DDD()
	default:
		DDD()
	}
}

var layers = []string{
	"domain",
	"application",
	"cmd",
	"di",
	"infra",
	"internal",
	"config",
}

func DDD() {
	// create root folder
	root := config.GlobalConfig.Project.Name
	if err := os.MkdirAll(root, 0o755); err != nil {
		panic(err)
	}

	// create ddd folder structure
	for _, layer := range layers {
		sub := filepath.Join(root, layer)
		if err := os.MkdirAll(sub, 0o755); err != nil {
			panic(err)
		}
	}
}
