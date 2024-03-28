/*
Copyright © 2023 Samuel Komfi <skomfi@gmail.com>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "md2pdf-cli",
	Short: "A go cli app that converts Markdown files to PDFs",
	Long: `The application uses the mdtopdf library to convert a Markdown file to a PDF file. For example:


This application is a tool to generate the needed pdf files
from a provided markdown file.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.md2pdf-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}


