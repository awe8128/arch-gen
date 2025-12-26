package config

type Config struct {
	Project PJ                `mapstructure:"pj"`
	Domains map[string]Domain `mapstructure:"domains"`
}

type PJ struct {
	Name string `mapstructure:"name"`
	Sys  string `mapstructure:"sys"`
}

type Domain struct {
	Type         string                `mapstructure:"type"`
	Properties   map[string]Property   `mapstructure:"properties"`
	Repositories map[string]Repository `mapstructure:"repositories"`
}

type Property struct {
	Type     string
	Nullable bool
}

type Repository struct {
	Type string
	In   map[string]Property
	Out  map[string]Property
}
