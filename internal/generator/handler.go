package generator

import (
	"fmt"

	"github.com/awe8128/arch-gen/shared/templates"
)

func HandlerTemplate() (string, string) {
	filename := "handler.go"

	template := fmt.Sprintf(
		`
		%s

		%s
		
		%s
		`,
		templates.NewPackageTemplate("server"),
		templates.HandlerStruct(),
		templates.NewHandlerFuncTemplate(),
	)

	return template, filename
}
