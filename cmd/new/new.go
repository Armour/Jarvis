package new

import (
	c "github.com/armour/jarvis/cmd"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "template",
	Short: "Start new project using templates.",
	Long:  "Start new project using templates.",
}

func init() {
	c.RootCmd.AddCommand(newCmd)
}
