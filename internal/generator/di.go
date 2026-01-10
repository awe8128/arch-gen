package generator

import (
	"fmt"
	"strings"
)

func NewDI(name string) (string, string) {
	filename := fmt.Sprintf("%s.go", name)
	template := fmt.Sprintf(`
	package di

	func %s(store *db.SQLStore) controller.%sController {
		repository := repository.New%sRepository(store)
		usecase := %s.New%sUsecase(repository)
		ctrl := controller.New%sController(usecase)

		return ctrl
	}
	`, capitalize(name), capitalize(name), capitalize(name), name, capitalize(name), capitalize(name))

	return template, filename
}

func capitalize(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}
