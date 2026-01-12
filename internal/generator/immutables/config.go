package immutables

import (
	"fmt"
	"path/filepath"

	"github.com/awe8128/arch-gen/templates"
	"github.com/awe8128/arch-gen/utils/fs"
)

func ConfigTemplate(root string) error {
	path := filepath.Join(root, "config")
	filename := "config.go"

	content := fmt.Sprintf(`
	%s
	%s

	func Load(path string) (config *Config, err error) {
	v := viper.New()

	v.AddConfigPath(path)
	v.SetConfigName(".app")
	v.SetConfigType("env")

	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		var notFound viper.ConfigFileNotFoundError
		if !errors.As(err, &notFound) {
			return nil, err
		}
	}

	// bind env keys so Unmarshal can see them even if file doesn't exist
	keys := []string{
		"DB_SOURCE",
		"DB_NAME",
		"DB_USER",
		"DB_PASSWORD",
		"DB_PORT",
		"DB_HOST",
		"MIGRATION_PATH",
		"API_PORT",
		"LOGGER_LEVEL",
	}
	for _, k := range keys {
		_ = v.BindEnv(k)
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
	func (c *Config) GetDSN() string {
	source := fmt.Sprintf("%s",
		c.DB_SOURCE,
		c.DB_USER,
		c.DB_PASSWORD,
		c.DB_HOST,
		c.DB_PORT,
		c.DB_NAME,
	)
	return source
}
	`,
		templates.Package("config"),
		templates.ConfigStruct(),
		"%s://%s:%s@%s:%s/%s?sslmode=disable",
	)
	if err := fs.GenerateFile(content, path, filename); err != nil {
		return err
	}
	return nil
}
