package generator

import (
	"fmt"

	"github.com/awe8128/arch-gen/config"
	"github.com/awe8128/arch-gen/templates"
)

// GenerateInfraRepository
func GenerateInfraRepository(pkg string, r map[string]config.Repository, p map[string]config.Property) (string, string) {
	filename := fmt.Sprintf("%s.go", pkg)

	content := fmt.Sprintf(`%s

	%s
	%s
	%s
	`,
		templates.Package("repository"),

		templates.Interface(pkg, "repository", pkg, r),

		templates.Store(pkg),

		templates.InterfaceMethod(pkg, r),
	)

	return content, filename
}
