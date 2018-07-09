// Package goproject contains GoCmd which generates new go project using template.
package goproject

import (
	"github.com/armour/jarvis/internal/pkg/config"
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
		Validate: survey.Required,
	},
	{
		Name: "projectDescription",
		Prompt: &survey.Input{
			Message: "Project Description?",
			Help:    "The description of the new project.",
		},
		Validate: survey.Required,
	},
	{
		Name: "license",
		Prompt: &survey.Select{
			Message: "Choose a license:",
			Options: []string{"MIT", "GPL", "Apache"},
			Default: "MIT",
		},
		Validate: survey.Required,
	},
}

// GoCmd generates new go project using template.
var GoCmd = &cobra.Command{
	Use:   "go",
	Short: "Start a new project using 'go' template",
	Long:  "Start a new project using 'go' template",
	Run: func(cmd *cobra.Command, args []string) {
		answers := struct {
			License            string
			ProjectDescription string
			ProjectName        string
		}{}
		if err := survey.Ask(questions, &answers); err != nil {
			utils.ExitOnError(err)
		}
		templates := []string{"new/go", "new/github", "new/license"}
		requireMap := map[string]interface{}{}
		replaceMap := map[string]interface{}{
			"author":             config.GetConfigByField("author"),
			"email":              config.GetConfigByField("email"),
			"githubUser":         config.GetConfigByField("githubUser"),
			"license":            answers.License,
			"projectDescription": answers.ProjectDescription,
			"projectName":        answers.ProjectName,
		}
		utils.GenerateFile(templates, "", requireMap, replaceMap)
	},
}
