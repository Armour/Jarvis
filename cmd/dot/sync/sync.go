// Package sync contains SyncCmd which downloads all global dot files to local machine.
package sync

import (
	"os/user"

	"github.com/armour/jarvis/internal/pkg/utils"
	"github.com/spf13/cobra"
	"gopkg.in/AlecAivazis/survey.v1"
)

var questions = []*survey.Question{
	{
		Name: "username",
		Prompt: &survey.Input{
			Message: "Username?",
			Default: "armour",
			Help:    "The name of current user.",
		},
	},
	{
		Name: "platform",
		Prompt: &survey.Select{
			Message: "Choose a platform:",
			Options: []string{"Mac", "Linux", "Windows"},
			Default: "Mac",
			Help:    "The platform for download dot files.",
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
			Username string
			Platform string
		}{}
		err := survey.Ask(questions, &answers)
		if err != nil {
			utils.ExitOnError(err)
		}

		templatePath := "../../../assets/dot/dotfiles"
		requireMap := map[string]interface{}{
			"mac":     answers.Platform == "Mac",
			"linux":   answers.Platform == "Linux",
			"windows": answers.Platform == "Windows",
		}
		replaceMap := map[string]interface{}{
			"username": answers.Username,
			"mac":      answers.Platform == "Mac",
			"linux":    answers.Platform == "Linux",
			"windows":  answers.Platform == "Windows",
		}
		usr, err := user.Current()
		if err != nil {
			utils.ExitOnError(err)
		}
		utils.GenerateFile(templatePath, usr.HomeDir, requireMap, replaceMap)
	},
}
