package main

import (
	"fmt"

	"github.com/markdown/markdownfmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "myapp",
	Short: "A simple command-line application",
	Long: `
    MyApp is a command-line tool that performs various tasks.

    For more information on specific commands, run:
    myapp help <command>
    `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running myapp...")
	},
}

func main() {
	// Add subcommands here

	// Set up usage templates
	rootCmd.SetUsageTemplate(markdownfmt.MustString(markdownfmt.Markdown(rootCmd.UsageTemplate)))
	rootCmd.SetHelpTemplate(markdownfmt.MustString(markdownfmt.Markdown(rootCmd.HelpTemplate)))

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
	}
}
