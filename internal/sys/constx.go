package sys

var (
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
		"schema",
		"server",
	}
)
