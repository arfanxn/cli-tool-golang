/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/arfanxn/cli-tool-golang/handlers"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cli-tool-golang",
	Short: "Get log file from the specified location and save it to the specified directory as a json or text file",
	Long:  `Get log file from the specified location and save it to the specified directory as a json or text file`,
	Run:   handlers.NewCMDHandler().RootHandler,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cli-tool-golang.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringP("output", "o", "", "Output file destination")
	rootCmd.Flags().StringP("type", "t", "text", "Format output as text or json")
}
