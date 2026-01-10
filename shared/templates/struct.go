package templates

import (
	"fmt"
	"strings"

	"github.com/awe8128/arch-gen/config"
)

func NewStructTemplate(name string, p map[string]config.Property) string {
	var fields strings.Builder

	for fieldName, prop := range p {
		fieldType := prop.Type
		if prop.Nullable {
			fieldType = "*" + fieldType
		}

		fields.WriteString(
			fmt.Sprintf("\t%s %s\n", capitalize(fieldName), fieldType),
		)
	}

	content := fmt.Sprintf(
		`
type %s struct {
%s}
`,
		capitalize(name),
		fields.String(),
	)

	return content
}

func capitalize(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func NewUsecaseStruct(name string) string {

	template := fmt.Sprintf(
		`
		type %sUsecase struct {
			repository repository.%sRepository
		}
		`, name, capitalize(name),
	)

	return template
}

func NewControllerStruct(name string) string {

	template := fmt.Sprintf(
		`
		type %sController struct {
			usecase %s.%sUsecase
		}
		`, name, name, capitalize(name),
	)

	return template
}

func NewDIfunc(name string) string {
	template := fmt.Sprintf(
		`
		func New%sUsecase(repository repository.%sRepository) %sUsecase {
			return &%sUsecase{
				repository: repository,
			}
		}
		`, capitalize(name), capitalize(name), capitalize(name), name,
	)

	return template

}

func NewDIfuncController(name string) string {
	template := fmt.Sprintf(
		`
		func New%sController(usecase %s.%sUsecase) %sController {
			return &%sController{
				usecase: usecase,
			}
		}
		`, capitalize(name), name, capitalize(name), capitalize(name), name,
	)

	return template

}

func HandlerStruct() string {
	var fields strings.Builder

	for name := range config.GlobalConfig.Domains {
		fields.WriteString(
			fmt.Sprintf("\t%s controller.%sController\n", capitalize(name), capitalize(name)),
		)
	}

	template := fmt.Sprintf(
		`
		type Handler struct {
			%s
		}

		`, fields.String(),
	)

	return template
}
