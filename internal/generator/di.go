package generator

import (
	"fmt"

	"github.com/awe8128/arch-gen/templates/utils"
)

func GenerateDI(name string) (string, string) {
	filename := fmt.Sprintf("%s.go", name)

	content := fmt.Sprintf(`
	package di

	func %s(store *db.SQLStore) controller.%sController {
		repository := repository.New%sRepository(store)
		usecase := %s.New%sUsecase(repository)
		ctrl := controller.New%sController(usecase)

		return ctrl
	}
	`,
		utils.Capitalize(name),
		utils.Capitalize(name),
		utils.Capitalize(name),
		name,
		utils.Capitalize(name),
		utils.Capitalize(name),
	)

	return content, filename
}
