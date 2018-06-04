// Package goproject contains GoCmd which generates new go project using template.
package goproject

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
			Help:    "The name of the new project, also will be used as go package name.",
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

// GoCmd generates new go project using template.
var GoCmd = &cobra.Command{
	Use:   "go",
	Short: "Start a new Go project",
	Long:  "Start a new Go project",
	Run: func(cmd *cobra.Command, args []string) {
		answers := struct {
			ProjectName string
			License     string
		}{}
		err := survey.Ask(questions, &answers)
		if err != nil {
			utils.ExitOnError(err)
		}

		templatePath := "../../../assets/new/go"
		requireMap := map[string]interface{}{}
		replaceMap := map[string]interface{}{
			"projectName": answers.ProjectName,
			"license":     answers.License,
		}
		utils.GenerateFile(templatePath, "", requireMap, replaceMap)
	},
}
