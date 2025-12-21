package config

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// configs
var (
	CfgFile string
)

type Config struct {
	Sys Sys `yaml:"sys"`
}

type Sys struct {
	Name string `yaml:"name"`
}

func Load() {

	if CfgFile != "" {
		viper.SetConfigFile(CfgFile)
	} else {
		fmt.Println("empty config string")
		// Find current directory.
		home, err := os.Getwd()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigName("arch")
		viper.SetConfigType("yaml")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(viper.ReadInConfig().Error())
	} else {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
