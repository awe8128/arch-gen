package builder

import (
	"fmt"
	"strings"

	"github.com/awe8128/arch-gen/config"
)

type structBuilder struct {
	name       string
	prefix     string
	suffix     string
	properties map[string]config.Property
}

func NewStructBuilder() *structBuilder {
	p := make(map[string]config.Property)
	return &structBuilder{
		properties: p,
	}
}

/*
typeName's are normal types
such as string, int, int64, bool etc.
*/
func (s *structBuilder) AddProperties(name string, typeName string, nullable bool) *structBuilder {
	p := config.Property{
		Type:     typeName,
		Nullable: nullable,
	}
	if _, ok := s.properties[name]; ok {
		panic("multiple struct body entry")
	}
	s.properties[name] = p
	return s
}

func (s *structBuilder) Name(name, prefix, suffix string) *structBuilder {

	s.name = name
	s.prefix = prefix
	s.suffix = suffix

	return s
}

func (s *structBuilder) Build() string {
	// struct Name
	structName := s.prefix + s.name + s.suffix

	// fields Name
	fields := buildField(s.properties)

	content := fmt.Sprintf(`
	type %s struct {
%s
}
`,
		structName, fields)
	return content
}

func buildField(properties map[string]config.Property) string {
	var fields strings.Builder

	for fieldName, prop := range properties {
		fieldType := prop.Type
		if prop.Nullable {
			fieldType = "*" + fieldType
		}

		fields.WriteString(
			fmt.Sprintf("\t%s %s\n", fieldName, fieldType),
		)
	}

	return fields.String()
}
