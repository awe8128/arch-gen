package generator

import (
	"fmt"

	"github.com/awe8128/arch-gen/config"
	"github.com/awe8128/arch-gen/shared/templates"
)

func ControllerTemplate(domain string, r map[string]config.Repository) (string, string) {
	filename := fmt.Sprintf(`%s.go`, domain)
	template := fmt.Sprintf(`
	%s

	//TODO: Add methods
	%s

	%s

	%s
	`,
		templates.NewPackageTemplate("controller"),
		templates.InterfaceTemplate(domain, "controller", nil),
		templates.NewControllerStruct(domain),
		templates.NewDIfuncController(domain),
	)
	return template, filename
}
