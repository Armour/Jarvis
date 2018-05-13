package cmd

import (
	"github.com/spf13/cobra"
)

var templateCmd = &cobra.Command{
	Use:   "template",
	Short: "Start new project using templates.",
	Long:  `Start new project using templates provided in 'templates' folder.`,
}

func init() {
	rootCmd.AddCommand(templateCmd)
}
