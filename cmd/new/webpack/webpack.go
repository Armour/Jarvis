// Package webpack contains WebpackCmd which generates new webpack project using template.
package webpack

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
			Help:    "The name of the new project, also will be used as npm package name.",
		},
	},
	{
		Name: "projectShortname",
		Prompt: &survey.Input{
			Message: "Project short name? (for PWA)",
			Default: "",
			Help:    "The description of the new project.",
		},
	},
	{
		Name: "projectDescription",
		Prompt: &survey.Input{
			Message: "Project description? (for PWA)",
			Default: "",
			Help:    "The description of the new project.",
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
	{
		Name: "jsFramework",
		Prompt: &survey.Select{
			Message: "Choose a javascript framework:",
			Options: []string{"React", "Vue"},
			Default: "React",
			Help:    "The javascript framework used in this project.",
		},
	},
	{
		Name: "cssFramework",
		Prompt: &survey.Select{
			Message: "Choose a css framework:",
			Options: []string{"Materialize"},
			Default: "Materialize",
			Help:    "The css framework used in this project.",
		},
	},
	{
		Name: "typescript",
		Prompt: &survey.Confirm{
			Message: "Use typescript?",
			Default: true,
		},
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
			Message: "Use CI?",
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

// WebpackCmd generates new webpack project using template.
var WebpackCmd = &cobra.Command{
	Use:   "webpack",
	Short: "Start a new project using 'webpack' template",
	Long:  "Start a new project using 'webpack' template",
	Run: func(cmd *cobra.Command, args []string) {
		answers := struct {
			ProjectName        string
			ProjectShortname   string
			ProjectDescription string
			License            string
			JsFramework        string
			CSSFramework       string
			Database           string
			CoverallToken      string
			Typescript         bool
			Redis              bool
			Docker             bool
			CI                 bool
		}{}
		err := survey.Ask(questions, &answers)
		if err != nil {
			utils.ExitOnError(err)
		}

		templatePath := "../../../assets/new/webpack"
		requireMap := map[string]interface{}{
			"ci":          answers.CI,
			"coverage":    answers.CoverallToken,
			"docker":      answers.Docker,
			"materialize": answers.CSSFramework == "Materialize",
			"postgres":    answers.Database == "Postgres",
			"react":       answers.JsFramework == "React",
			"redis":       answers.Redis,
			"typescript":  answers.Typescript,
		}
		if answers.ProjectShortname == "" {
			answers.ProjectShortname = answers.ProjectName
		}
		replaceMap := map[string]interface{}{
			"ci":                 answers.CI,
			"coverage":           answers.CoverallToken,
			"docker":             answers.Docker,
			"materialize":        answers.CSSFramework == "Materialize",
			"postgres":           answers.Database == "Postgres",
			"react":              answers.JsFramework == "React",
			"redis":              answers.Redis,
			"typescript":         answers.Typescript,
			"projectName":        answers.ProjectName,
			"projectShortname":   answers.ProjectShortname,
			"projectDescription": answers.ProjectDescription,
			"license":            answers.License,
		}
		utils.GenerateFile(templatePath, "", requireMap, replaceMap)
	},
}
