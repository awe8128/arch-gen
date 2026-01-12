package generator

import (
	"fmt"

	"github.com/awe8128/arch-gen/config"
	"github.com/awe8128/arch-gen/templates"
)

func GenerateEntity(name string, p map[string]config.Property) (string, string) {
	filename := fmt.Sprintf("%s.go", name)

	content := fmt.Sprintf(`%s

	%s
	%s
	`,
		templates.Package(name),
		templates.NewStruct("", name, "Entity", p),
		templates.NewEntityFunc(name, p, nil),
	)

	return content, filename
}
