package cmd

import (
	"github.com/armour/jarvis/utils"
	"github.com/spf13/cobra"
)

var miscCmd = &cobra.Command{
	Use:   "misc",
	Short: "Start new project using misc template",
	Long:  "Start new project using misc template",
	Run: func(cmd *cobra.Command, args []string) {
		templatePath := "../new/misc"
		requireMap := map[string]interface{}{}
		replaceMap := map[string]interface{}{
			"projectName": projectName,
			"license":     license,
		}
		utils.GenerateFile(templatePath, "", requireMap, replaceMap)
	},
}

func init() {
	newCmd.AddCommand(miscCmd)
	miscCmd.Flags().StringVarP(&projectName, "projectName", "n", "", "The name for this project. (required)")
	miscCmd.Flags().StringVarP(&license, "license", "l", "MIT", "The license for this project, default is \"MIT\". (\"MIT\", \"GPL\", \"Apache\"")

	miscCmd.MarkFlagRequired("projectName")
}
