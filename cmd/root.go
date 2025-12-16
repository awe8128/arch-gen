/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/awe8128/arch-gen/cmd/folder"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// configs
var (
	cfgFile string
)

// rootCmd represents the base command when called without any subcommands
var (
	rootCmd = &cobra.Command{
		Use:   "arch-gen",
		Short: "Generative tool for Backend System Design",
		Long: `Generate .go files and folder structure for specific system design
for example:
	Domain Driven Architecture,
	Layer Architecture 
		`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello world")
			fmt.Println(viper.Get("sys"))
			fmt.Println(viper.Get("sys.name"))
		},
	}
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// config + flags mainly
func init() {
	cobra.OnInitialize(LoadConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/arch-gen.yaml)")

	rootCmd.AddCommand(folder.GenerateFolderCmd)
}

type Config struct {
	Sys Sys `yaml:"sys"`
}

type Sys struct {
	Name string `yaml:"name"`
}

func LoadConfig() {

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
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
