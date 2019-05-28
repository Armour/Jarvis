// Package sync contains SyncCmd which downloads all global dot files to local machine.
package sync

import (
	"os/user"

	"github.com/armour/jarvis/internal/pkg/utils"
	"github.com/spf13/cobra"
	survey "gopkg.in/AlecAivazis/survey.v1"
)

var questions = []*survey.Question{
	{
		Name: "platform",
		Prompt: &survey.Select{
			Message: "Choose a platform:",
			Options: []string{"Mac", "Linux", "Windows"},
			Default: "Mac",
			Help:    "The platform of your machine.",
		},
	},
}

// SyncCmd downloads all global dot files to local machine.
var SyncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Download all global dot files to local machine.",
	Long:  "Download all global dot files to local machine.",
	Run: func(cmd *cobra.Command, args []string) {
		answers := struct {
			Platform string
		}{}
		if err := survey.Ask(questions, &answers); err != nil {
			utils.ExitOnError(err)
		}
		usr, err := user.Current()
		if err != nil {
			utils.ExitOnError(err)
		}
		templates := []string{"dot/dotfiles"}
		requireMap := map[string]interface{}{
			"mac":     answers.Platform == "Mac",
			"linux":   answers.Platform == "Linux",
			"windows": answers.Platform == "Windows",
		}
		replaceMap := map[string]interface{}{
			"username": usr.Username,
			"mac":      answers.Platform == "Mac",
			"linux":    answers.Platform == "Linux",
			"windows":  answers.Platform == "Windows",
		}
		utils.GenerateFile(templates, usr.HomeDir, requireMap, replaceMap)
	},
}
