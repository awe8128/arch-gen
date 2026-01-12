package generator

import (
	"fmt"

	"github.com/awe8128/arch-gen/templates"
)

func GenerateHandler() (string, string) {
	filename := "handler.go"

	content := fmt.Sprintf(
		`
		%s

		%s
		
		%s
		`,
		templates.Package("server"),
		templates.HandlerStruct(),
		templates.NewHandlerFunc(),
	)

	return content, filename
}
