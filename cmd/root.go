/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/awe8128/arch-gen/cmd/folder"
	"github.com/awe8128/arch-gen/config"
	"github.com/spf13/cobra"
)

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

			config.Get("root2")
		},
	}
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// root
// TODO: load config default or custom yaml
// TODO: fetch all data convert to custom type

// TODO: make sure we can use loaded config from any level
// ex: config.Get("sys") = map[string]string of sys

func init() {
	cobra.OnInitialize(config.Load)

	rootCmd.PersistentFlags().StringVarP(&config.CfgFile, "config", "c", "", "config file (default is $HOME/arch-gen.yaml)")

	rootCmd.AddCommand(folder.GenerateFolderCmd)
}
