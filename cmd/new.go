package cmd

import (
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Start new project using templates.",
	Long:  "Start new project using templates.",
}

func init() {
	rootCmd.AddCommand(newCmd)
}
