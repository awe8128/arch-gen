package generator

import (
	"fmt"

	"github.com/awe8128/arch-gen/shared/templates"
	"github.com/awe8128/arch-gen/templates/immutable"
)

func HandlerTemplate() (string, string) {
	filename := "handler.go"

	template := fmt.Sprintf(
		`
		%s

		%s
		
		%s
		`,
		templates.PackageTemplate("server"),
		templates.HandlerStruct(),
		immutable.NewHandlerFuncTemplate(),
	)

	return template, filename
}
