package cmd

import (
	"github.com/armour/jarvis/utils"
	"github.com/spf13/cobra"
	"gopkg.in/AlecAivazis/survey.v1"
)

var miscQuestions = []*survey.Question{
	{
		Name: "projectName",
		Prompt: &survey.Input{
			Message: "Project name?",
			Default: "my-project",
			Help:    "The name of the new project, also will be used as npm package name.",
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

var miscCmd = &cobra.Command{
	Use:   "misc",
	Short: "Start new project using misc template",
	Long:  "Start new project using misc template",
	Run: func(cmd *cobra.Command, args []string) {
		miscAnswers := struct {
			ProjectName string
			License     string
		}{}
		err := survey.Ask(miscQuestions, &miscAnswers)
		if err != nil {
			utils.ExitOnError(err)
		}

		templatePath := "../new/misc"
		requireMap := map[string]interface{}{}
		replaceMap := map[string]interface{}{
			"projectName": miscAnswers.ProjectName,
			"license":     miscAnswers.License,
		}
		utils.GenerateFile(templatePath, "", requireMap, replaceMap)
	},
}

func init() {
	newCmd.AddCommand(miscCmd)
}
