package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "jarvis",
	Short: "My personal assistant.",
	Long:  "Jarvis is my personal assistant who can help me do some tedious work.",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		exitOnError(err)
	}
}
