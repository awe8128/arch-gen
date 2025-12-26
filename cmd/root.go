/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/awe8128/arch-gen/cmd/generate"
	"github.com/awe8128/arch-gen/config"
	"github.com/awe8128/arch-gen/internal/linter"
	"github.com/awe8128/arch-gen/internal/structure"
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
			structure.Start()
			linter.FormatAndFixImportsDir("./be")
		},
	}
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(config.Load)

	rootCmd.PersistentFlags().StringVarP(&config.CfgFile, "config", "c", "", "config file (default is $HOME/arch-gen.yaml)")

	rootCmd.AddCommand(generate.GenerateFolderCmd)
}
