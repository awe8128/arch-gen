package templates

import "fmt"

func NewPackageTemplate(pkg string) string {

	template := fmt.Sprintf(`package %s`, pkg)

	return template
}
