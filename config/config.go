package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// configs
var (
	CfgFile string
)

var GlobalConfig *Config

func Load() {

	v := viper.New()

	if CfgFile != "" {
		v.SetConfigFile(CfgFile)
	} else {
		fmt.Println("empty config string")

		// Find current directory.
		home, err := os.Getwd()
		cobra.CheckErr(err)

		v.AddConfigPath(home)
		v.SetConfigName("arch")
		v.SetConfigType("yaml")
	}

	if err := v.ReadInConfig(); err != nil {
		fmt.Println("Error reading config:", err)
	}

	fmt.Println("Using config file:", v.ConfigFileUsed())

	if err := v.Unmarshal(&GlobalConfig); err != nil {
		log.Fatal("Failed to load config", err)
	}

}
