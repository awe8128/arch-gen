package templates

import "fmt"

func PackageTemplate(name string) string {

	template := fmt.Sprintf(`package %s`, name)

	return template
}
