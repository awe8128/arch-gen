package filesx

import (
	"fmt"

	"github.com/spf13/viper"
)

type Yaml struct{}

func ReaderYaml() {
	d := viper.AllSettings()
	fmt.Println(d)
}
