package templates

import "fmt"

func NewPackageTemplate(name string) string {

	template := fmt.Sprintf(`package %s`, name)

	return template
}
