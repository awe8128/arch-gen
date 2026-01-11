package immutable

import (
	"fmt"
	"sort"
	"strings"

	"github.com/awe8128/arch-gen/config"
)

type funcBuilder struct {
	name   string
	prefix string
	in     map[string]config.Property
	out    map[string]config.Property
	method string
	body   string
}

func NewFuncBuilder(prefix, funcName string, in, out map[string]config.Property) *funcBuilder {
	return &funcBuilder{
		name:   capitalize(funcName),
		prefix: capitalize(prefix),
		out:    out,
		in:     in,
	}
}

func (f *funcBuilder) BuildFunc(isCtx bool) string {
	content := fmt.Sprintf(`func %s %s %s %s %s`,
		f.method,
		f.prefix+f.name,
		GetParams(f.in, isCtx),
		GetReturnValues(f.name, f.out),
		f.body,
	)

	return content
}

func (f *funcBuilder) BuildInterface(isCtx bool) string {
	content := fmt.Sprintf(`%s %s %s %s %s`,
		f.method,
		f.prefix+f.name,
		GetParams(f.in, isCtx),
		GetReturnValues(f.name, f.out),
		f.body,
	)

	return content
}

func (f *funcBuilder) Method(m string) *funcBuilder {

	content := fmt.Sprintf(`(%s *%s)`, strings.ToLower(string(m[0])), m)

	f.method = content

	return f
}

func (f *funcBuilder) Body() *funcBuilder {
	var r strings.Builder
	var res string

	// Sort keys for consistent ordering
	keys := make([]string, 0, len(f.out))
	for k := range f.out {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key := range keys {
		prop := f.out[key]
		fieldType := prop.Type

		switch fieldType {
		case "uuid.UUID":
			res = "uuid.New()"
		case "string":
			res = ""
		case "int", "int64", "uint", "int32":
			res = fmt.Sprintf("%s(0)", fieldType)
		case "bool":
			res = "false"
		default:
			res = ""
		}

		if prop.Nullable {
			res = "nil"
		}

		r.WriteString(
			fmt.Sprintf("%s,", res),
		)
	}

	content := fmt.Sprintf(`{
		return %s
	}`, r.String()[:len(r.String())-1])

	f.body = content
	return f
}

// Specific usecase
func NewEntityFuncTemplate(name string, p, r map[string]config.Property) string {

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
			fmt.Sprintf("\t%s:%s,\n", capitalize(key), key),
		)

	}

	fn := NewFuncBuilder("New", capitalize(name), p, nil).BuildFunc(true)

	contextContent := fmt.Sprintf(`
	ent:=&%s{
		%s
	}
	
	`, GetReturnValues(name, r)[1:], funcProcess.String())

	template = fmt.Sprintf(`
%s{
	%s
	return ent
}
	`, fn, contextContent)

	return template
}

func NewHandlerFuncTemplate() string {

	var fields strings.Builder

	// Sort keys for consistent ordering
	keys := make([]string, 0, len(config.GlobalConfig.Domains))
	for k := range config.GlobalConfig.Domains {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, name := range keys {
		fields.WriteString(
			fmt.Sprintf("\t%s: di.%s(db),\n", capitalize(name), capitalize(name)),
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
