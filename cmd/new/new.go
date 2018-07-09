// Package new contains NewCmd which generates new project using templates.
package new

import (
	"github.com/armour/jarvis/cmd/new/go"
	"github.com/armour/jarvis/cmd/new/misc"
	"github.com/armour/jarvis/cmd/new/python"
	"github.com/armour/jarvis/cmd/new/react"
	"github.com/armour/jarvis/cmd/new/unity"
	"github.com/spf13/cobra"
)

// NewCmd generates new project using templates.
var NewCmd = &cobra.Command{
	Use:   "new",
	Short: "Start new project using templates",
	Long:  "Start new project using templates",
}

func init() {
	NewCmd.AddCommand(goproject.GoCmd)
	NewCmd.AddCommand(misc.MiscCmd)
	NewCmd.AddCommand(python.PythonCmd)
	NewCmd.AddCommand(unity.UnityCmd)
	NewCmd.AddCommand(react.ReactCmd)
}
