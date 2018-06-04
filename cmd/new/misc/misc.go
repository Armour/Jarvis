// Package misc contains MiscCmd which generates new misc project using template.
package misc

import (
	"github.com/armour/jarvis/internal/pkg/utils"
	"github.com/spf13/cobra"
	"gopkg.in/AlecAivazis/survey.v1"
)

var questions = []*survey.Question{
	{
		Name: "projectName",
		Prompt: &survey.Input{
			Message: "Project name?",
			Default: "my-project",
			Help:    "The name of the new project.",
		},
	},
	{
		Name: "license",
		Prompt: &survey.Select{
			Message: "Choose a license:",
			Options: []string{"MIT", "GPL", "Apache"},
			Default: "MIT",
			Help:    "The license of this project.",
		},
	},
}

// MiscCmd generates new misc project using template.
var MiscCmd = &cobra.Command{
	Use:   "misc",
	Short: "Start new project using misc template",
	Long:  "Start new project using misc template",
	Run: func(cmd *cobra.Command, args []string) {
		answers := struct {
			ProjectName string
			License     string
		}{}
		err := survey.Ask(questions, &answers)
		if err != nil {
			utils.ExitOnError(err)
		}

		templatePath := "../../../assets/new/misc"
		requireMap := map[string]interface{}{}
		replaceMap := map[string]interface{}{
			"projectName": answers.ProjectName,
			"license":     answers.License,
		}
		utils.GenerateFile(templatePath, "", requireMap, replaceMap)
	},
}
