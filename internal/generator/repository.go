package generator

import (
	"fmt"

	"github.com/awe8128/arch-gen/config"
	"github.com/awe8128/arch-gen/templates"
)

func GenerateRepository(pkg string, r map[string]config.Repository) (string, string) {
	filename := fmt.Sprintf("%s.go", "repository")
	content := fmt.Sprintf(`%s

	%s
	`,
		templates.Package(pkg),
		templates.Interface("repository", "", pkg, r),
	)

	return content, filename
}
