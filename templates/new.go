package templates

import (
	"fmt"
	"sort"
	"strings"

	"github.com/awe8128/arch-gen/config"
	"github.com/awe8128/arch-gen/templates/builder"
	"github.com/awe8128/arch-gen/templates/utils"
)

func NewEntityFunc(name string, p, r map[string]config.Property) string {

	var template string
	var funcProcess strings.Builder

	// Sort keys for consistent ordering
	keys := make([]string, 0, len(p))
	for k := range p {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key := range keys {
		funcProcess.WriteString(
			fmt.Sprintf("\t%s:%s,\n", utils.Capitalize(key), key),
		)

	}
	fn := builder.NewFuncBuilder().
		Name("New", utils.Capitalize(name), "").
		In(p).
		AddOutProperty("", fmt.Sprintf("%sEntity", utils.Capitalize(name)), true).
		BuildFunc(false)

	contextContent := fmt.Sprintf(`
	ent:=&%sEntity{
		%s
	}
	
	`, utils.GetReturnValues(name, r)[1:], funcProcess.String())

	template = fmt.Sprintf(`
%s{
	%s
	return ent
}
	`, fn, contextContent)

	return template
}

func NewHandlerFunc() string {

	var fields strings.Builder

	// Sort keys for consistent ordering
	keys := make([]string, 0, len(config.GlobalConfig.Domains))
	for k := range config.GlobalConfig.Domains {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, name := range keys {
		fields.WriteString(
			fmt.Sprintf("\t%s: di.%s(db),\n", utils.Capitalize(name), utils.Capitalize(name)),
		)
	}

	template := fmt.Sprintf(
		`
			func NewHandler(db *db.SQLStore) *Handler {
				return &Handler{
					%s
				}
			}
	
		`, fields.String(),
	)

	return template
}
