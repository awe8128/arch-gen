package templates

import "fmt"

func Package(name string) string {

	template := fmt.Sprintf(`package %s`, name)

	return template
}
