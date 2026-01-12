package builder

import (
	"fmt"
	"sort"
	"strings"

	"github.com/awe8128/arch-gen/config"
	"github.com/awe8128/arch-gen/templates/utils"
)

type funcBuilder struct {
	name   string
	prefix string
	suffix string
	in     map[string]config.Property
	out    map[string]config.Property
	method string
	body   string
}

func NewFuncBuilder() *funcBuilder {
	in := make(map[string]config.Property)
	out := make(map[string]config.Property)

	return &funcBuilder{
		in:  in,
		out: out,
	}
}

func (f *funcBuilder) Name(prefix, name, suffix string) *funcBuilder {

	f.name = name
	f.prefix = prefix
	f.suffix = suffix

	return f
}
func (f *funcBuilder) In(in map[string]config.Property) *funcBuilder {
	if in != nil {
		f.in = in
	}

	return f
}

func (f *funcBuilder) AddInProperty(name string, typeName string, nullable bool) *funcBuilder {
	p := config.Property{
		Type:     typeName,
		Nullable: nullable,
	}

	if _, ok := f.in[name]; ok {
		panic("multiple struct body entry")
	}

	f.in[name] = p

	return f
}

func (f *funcBuilder) Out(out map[string]config.Property) *funcBuilder {
	if out != nil {
		f.out = out

	}
	return f
}

func (f *funcBuilder) AddOutProperty(name string, typeName string, nullable bool) *funcBuilder {
	p := config.Property{
		Type:     typeName,
		Nullable: nullable,
	}

	if _, ok := f.out[name]; ok {
		panic("multiple struct body entry")
	}
	f.out[name] = p
	return f
}

func (f *funcBuilder) BuildFunc(isCtx bool) string {
	name := f.prefix + f.name + f.suffix
	content := fmt.Sprintf(`func %s %s %s %s %s`,
		f.method,
		name,
		utils.GetParams(f.in, isCtx),
		utils.GetReturnValues(f.name, f.out),
		f.body,
	)

	return content
}

func (f *funcBuilder) BuildInterface(isCtx bool) string {
	name := f.prefix + f.name + f.suffix
	content := fmt.Sprintf(`%s %s %s %s %s`,
		f.method,
		name,
		utils.GetParams(f.in, isCtx),
		utils.GetReturnValues(f.name, f.out),
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
			if strings.Contains(fieldType, "Usecase") {
				res = newFunc(prop.Type, fieldType)
			}
			if strings.Contains(fieldType, "Controller") {
				res = newFunc(prop.Type, fieldType)
			}
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

func newFunc(field, fieldType string) string {
	var res string
	if strings.Contains(fieldType, "Usecase") {
		res = fmt.Sprintf(`%s{
			repository: repository,
		}`, strings.ToLower(string(field[0]))+string(field[1:]))
	}

	if strings.Contains(fieldType, "Controller") {
		res = fmt.Sprintf(`%s{
			usecase: usecase,
		}`, strings.ToLower(string(field[0]))+string(field[1:]))
	}

	return res
}
