package generator

import (
	"fmt"

	"github.com/awe8128/arch-gen/config"
	"github.com/awe8128/arch-gen/templates"
)

func EntityTemplate(name string, p map[string]config.Property) (string, string) {
	filename := fmt.Sprintf("%s.go", name)

	content := fmt.Sprintf(`%s

	%s
	%s
	`,
		templates.PackageTemplate(name),
		templates.NewStructTemplate("", name, "Entity", p),
		templates.NewEntityFuncTemplate(name, p, nil),
	)

	return content, filename
}
