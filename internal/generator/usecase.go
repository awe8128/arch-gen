package generator

import (
	"fmt"

	"github.com/awe8128/arch-gen/config"
	"github.com/awe8128/arch-gen/shared/templates"
)

func UsecaseTemplate(domain string, r map[string]config.Repository) (string, string) {
	filename := fmt.Sprintf(`%s.go`, domain)
	template := fmt.Sprintf(`
	%s

	//TODO: Add methods
	%s

	%s

	%s
	`,
		templates.NewPackageTemplate(domain),
		templates.InterfaceTemplate(domain, "usecase", nil),
		templates.NewUsecaseStruct(domain),
		templates.NewDIfunc(domain),
	)
	return template, filename
}
