// Package python contains PythonCmd which generates new python project using template.
package python

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

// PythonCmd generates new python project using template.
var PythonCmd = &cobra.Command{
	Use:   "python",
	Short: "Start a new project using 'python' template",
	Long:  "Start a new project using 'python' template",
	Run: func(cmd *cobra.Command, args []string) {
		answers := struct {
			ProjectName string
			License     string
		}{}
		err := survey.Ask(questions, &answers)
		if err != nil {
			utils.ExitOnError(err)
		}

		templatePath := "../../../assets/new/python"
		requireMap := map[string]interface{}{}
		replaceMap := map[string]interface{}{
			"projectName": answers.ProjectName,
			"license":     answers.License,
		}
		utils.GenerateFile(templatePath, "", requireMap, replaceMap)
	},
}
