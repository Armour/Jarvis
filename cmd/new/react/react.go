// Package react contains ReactCmd which generates new react project using template.
package react

import (
	"fmt"

	"github.com/armour/jarvis/internal/pkg/config"
	"github.com/armour/jarvis/internal/pkg/utils"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	survey "gopkg.in/AlecAivazis/survey.v1"
)

var questions = []*survey.Question{
	{
		Name: "projectName",
		Prompt: &survey.Input{
			Message: "Project name?",
			Default: "my-project",
			Help:    "The name of the new project, also will be used as npm package name.",
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
	{
		Name: "database",
		Prompt: &survey.Select{
			Message: "Choose a database:",
			Options: []string{"Postgres"},
			Default: "Postgres",
			Help:    "The database used in this project.",
		},
	},
	{
		Name: "redis",
		Prompt: &survey.Confirm{
			Message: "Use redis as cache store?",
			Default: true,
		},
	},
	{
		Name: "docker",
		Prompt: &survey.Confirm{
			Message: "Use docker?",
			Default: true,
		},
	},
	{
		Name: "ci",
		Prompt: &survey.Confirm{
			Message: "Use CircleCI?",
			Default: true,
		},
	},
	{
		Name: "coverallToken",
		Prompt: &survey.Input{
			Message: "CoverallToken? (for code coverage)",
			Default: "",
			Help:    "The token for using coverall.",
		},
	},
}

// ReactCmd generates new react project using template.
var ReactCmd = &cobra.Command{
	Use:   "react",
	Short: "Start a new project using 'react' template",
	Long:  "Start a new project using 'react' template",
	Run: func(cmd *cobra.Command, args []string) {
		answers := struct {
			CI                 bool
			CoverallToken      string
			Database           string
			Docker             bool
			License            string
			ProjectDescription string
			ProjectName        string
			Redis              bool
		}{}
		if err := survey.Ask(questions, &answers); err != nil {
			utils.ExitOnError(err)
		}
		templates := []string{"new/react", "new/github", "new/license"}
		requireMap := map[string]interface{}{
			"ci":       answers.CI,
			"coverage": answers.CoverallToken,
			"docker":   answers.Docker,
			"postgres": answers.Database == "Postgres",
			"redis":    answers.Redis,
		}
		replaceMap := map[string]interface{}{
			"author":             config.GetConfigByField("author"),
			"ci":                 answers.CI,
			"coverallToken":      answers.CoverallToken,
			"docker":             answers.Docker,
			"dockerUser":         config.GetConfigByField("dockerUser"),
			"email":              config.GetConfigByField("email"),
			"githubUser":         config.GetConfigByField("githubUser"),
			"license":            answers.License,
			"postgres":           answers.Database == "Postgres",
			"projectDescription": answers.ProjectDescription,
			"projectName":        answers.ProjectName,
			"redis":              answers.Redis,
		}
		utils.GenerateFile(templates, answers.ProjectName, requireMap, replaceMap)

		bold := color.New(color.Bold).SprintFunc()
		boldCyan := color.New(color.FgCyan, color.Bold).SprintFunc()

		fmt.Println()
		fmt.Printf("%s%s%s\n", bold("This project uses "), boldCyan("Commitlint"), bold(", see more detail here: https://github.com/Armour/commitlint-config-armour"))
	},
}
