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

var (
	layers = []string{
		"domain",
		"application",
		"cmd",
		"di",
		"infra",
		"internal",
		"config",
		"openapi",
		"presentation",
	}

	openapiLayers = []string{
		"components",
		"paths",
	}

	openapiComponents = []string{
		"requestBodies",
		"responses",
		"schemas",
	}

	serverLayer = []string{
		"logx",
		"middleware",
	}

	applicationLayer = []string{
		"usecase",
	}

	infrastructureLayer = []string{
		"db",
		"repository",
	}

	DBLayer = []string{
		"migrations",
		"query",
	}

	PresentationLayer = []string{
		"controller",
		"openapi",
		"server",
	}

	cmdLayer = []string{
		"api",
		"migrate",
	}
)

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
	// openapi/**
	for _, layer := range openapiLayers {
		sub := filepath.Join(root, "openapi", layer)
		if err := os.MkdirAll(sub, 0o755); err != nil {
			panic(err)
		}
	}

	// openapi/components/**
	for _, layer := range openapiComponents {
		sub := filepath.Join(root, "openapi", "components", layer)
		if err := os.MkdirAll(sub, 0o755); err != nil {
			panic(err)
		}
	}

	for _, layer := range serverLayer {
		sub := filepath.Join(root, "presentation", "server", layer)
		if err := os.MkdirAll(sub, 0o755); err != nil {
			panic(err)
		}
	}

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

	for _, folder := range PresentationLayer {
		path := filepath.Join(root, "presentation", folder)
		if err := os.MkdirAll(path, 0o755); err != nil {
			panic(err)
		}
	}

	for _, folder := range cmdLayer {
		path := filepath.Join(root, "cmd", folder)
		if err := os.MkdirAll(path, 0o755); err != nil {
			panic(err)
		}
	}
}
