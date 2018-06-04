// Package cmd contains rootCmd which is the command line entry for Jarvis.
package cmd

import (
	"github.com/armour/jarvis/cmd/dot"
	"github.com/armour/jarvis/cmd/new"
	"github.com/armour/jarvis/cmd/say"
	"github.com/armour/jarvis/internal/pkg/utils"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "jarvis",
	Short: "My personal assistant",
	Long:  "Jarvis is my personal assistant who can help me do some tedious work",
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		utils.ExitOnError(err)
	}
}

func init() {
	rootCmd.AddCommand(dot.DotCmd)
	rootCmd.AddCommand(new.NewCmd)
	rootCmd.AddCommand(say.SayCmd)
}
