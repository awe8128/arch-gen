package generator

import (
	"fmt"

	"github.com/awe8128/arch-gen/templates"
)

func GenerateHandler() (string, string) {
	filename := "handler.go"

	template := fmt.Sprintf(
		`
		%s

		%s
		
		%s
		`,
		templates.PackageTemplate("server"),
		templates.HandlerStructTemplate(),
		templates.NewHandlerFuncTemplate(),
	)

	return template, filename
}
