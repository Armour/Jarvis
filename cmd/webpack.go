package cmd

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/gobuffalo/packr"
	"github.com/spf13/cobra"
)

var (
	coverallToken      string
	projectName        string
	projectShortname   string
	projectDescription string
	projectBgColor     string
	projectThemeColor  string
	ci                 bool
	docker             bool
	materialize        bool
	postgres           bool
	react              bool
	redis              bool
	typescript         bool
)

const (
	requirePattern = `\[(.*)\]`
)

var (
	requireRE = regexp.MustCompile(requirePattern)
)

var webpackCmd = &cobra.Command{
	Use:   "webpack",
	Short: "Start new project using webpack template.",
	Long:  `Start new project using webpack template provided in 'templates/webpack' folder.`,
	Run: func(cmd *cobra.Command, args []string) {
		requireMap := map[string]interface{}{
			"ci":          ci,
			"coverage":    coverallToken,
			"docker":      docker,
			"materialize": materialize,
			"postgres":    postgres,
			"react":       react,
			"redis":       redis,
			"typescript":  typescript,
		}
		if projectShortname == "" {
			projectShortname = projectName
		}
		fmt.Println("Generated file:")
		templatesBox := packr.NewBox("../templates/webpack")
		for _, f := range templatesBox.List() {
			var missRequired bool
			outputFile := f
			requireMatches := requireRE.FindAllStringSubmatch(f, -1)
			for _, m := range requireMatches {
				if len(m) > 1 {
					for _, word := range m {
						if r, ok := requireMap[word]; ok {
							if r == false || r == "" {
								missRequired = true
								break
							}
						}
					}
					outputFile = strings.Replace(outputFile, m[0], "", -1)
				}
				if missRequired {
					break
				}
			}
			if missRequired {
				continue
			}
			content := templatesBox.String(f)
			t := template.Must(template.New("template").Parse(content))
			var tpl bytes.Buffer
			err := t.Execute(&tpl, map[string]interface{}{
				"ci":                 ci,
				"coverallToken":      coverallToken,
				"docker":             docker,
				"materialize":        materialize,
				"postgres":           postgres,
				"projectName":        projectName,
				"projectShortname":   projectShortname,
				"projectDescription": projectDescription,
				"projectBgColor":     projectBgColor,
				"projectThemeColor":  projectThemeColor,
				"react":              react,
				"redis":              redis,
				"typescript":         typescript,
			})
			if err != nil {
				exitOnError(err)
			}
			outputFile = strings.TrimSuffix(outputFile, ".gotmpl")
			folderIndex := strings.LastIndex(outputFile, "/")
			if folderIndex != -1 {
				folderPath := outputFile[:folderIndex]
				if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
					exitOnError(err)
				}
			}
			err = ioutil.WriteFile(outputFile, tpl.Bytes(), 0644)
			exitOnError(err)
			fmt.Println(outputFile)
		}
	},
}

func init() {
	templateCmd.AddCommand(webpackCmd)

	webpackCmd.Flags().StringVarP(&coverallToken, "coverallToken", "c", "", "The token for coverall code coverage. (optional)")
	webpackCmd.Flags().StringVarP(&projectName, "projectName", "n", "", "The name for this new project. (required)")
	webpackCmd.Flags().StringVarP(&projectShortname, "projectShortname", "s", "", "The short name for this new project. (optional)")
	webpackCmd.Flags().StringVarP(&projectDescription, "projectDescription", "d", "project description here", "The description for this new project. (optional)")
	webpackCmd.Flags().StringVarP(&projectBgColor, "projectBgColor", "b", "#2196f3", "The background color for this new project. (optional)")
	webpackCmd.Flags().StringVarP(&projectThemeColor, "projectThemeColor", "t", "#2196f3", "The theme color for this new project. (optional)")
	webpackCmd.Flags().BoolVar(&ci, "ci", true, "The flag to enable continuous integration support.")
	webpackCmd.Flags().BoolVar(&docker, "docker", true, "The flag to enable docker support.")
	webpackCmd.Flags().BoolVar(&materialize, "materialize", true, "The flag to enable materialize support.")
	webpackCmd.Flags().BoolVar(&postgres, "postgres", true, "The flag to enable postgres support.")
	webpackCmd.Flags().BoolVar(&react, "react", true, "The flag to enable react support.")
	webpackCmd.Flags().BoolVar(&redis, "redis", true, "The flag to enable redis support.")
	webpackCmd.Flags().BoolVar(&typescript, "typescript", true, "The flag to enable typescript support.")

	webpackCmd.MarkFlagRequired("projectName")
}
