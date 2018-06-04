// Package dot contains DotCmd which manages global dot files.
package dot

import (
	"github.com/armour/jarvis/cmd/dot/sync"
	"github.com/spf13/cobra"
)

// DotCmd manages global dot files.
var DotCmd = &cobra.Command{
	Use:   "dot",
	Short: "Manage global dot files",
	Long:  "Manage global dot files",
}

func init() {
	DotCmd.AddCommand(sync.SyncCmd)
}
