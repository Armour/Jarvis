// Package utils contains all utility functions that will be used across Jarvis.
package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

	"github.com/fatih/color"
	"github.com/gobuffalo/packr"
)

const (
	requirePattern = `\[(.*)\]`
)

var (
	requireRE = regexp.MustCompile(requirePattern)
)

// ExitOnError prints the error message and exit with code 1.
func ExitOnError(err error) {
	if err != nil {
		color.Red("%q", err)
		os.Exit(1)
	}
}

// RunCommand executes input command and prints the result.
func RunCommand(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	var outBuf bytes.Buffer
	var errBuf bytes.Buffer
	cmd.Stdout = &outBuf
	cmd.Stderr = &errBuf
	if err := cmd.Run(); err != nil {
		if errBuf.String() != "" {
			color.Red("%q", errBuf.String())
		}
		ExitOnError(err)
	}
	if outBuf.String() != "" {
		fmt.Printf("%q", outBuf.String())
	}
}

// GenerateFile generates files using template in assets folder.
func GenerateFile(templates []string, outputPath string, requireMap map[string]interface{}, replaceMap map[string]interface{}) {
	assetsPath := "../../../assets/"
	for _, t := range templates {
		templatesBox := packr.NewBox(assetsPath + t)
		bold := color.New(color.Bold).SprintFunc()
		boldGreen := color.New(color.FgGreen, color.Bold).SprintFunc()
		fmt.Printf("\n%s%s%s\n", bold("Using template "), boldGreen(t), bold(", generated file:"))
		for _, f := range templatesBox.List() {
			// Check requirement.
			var invalid bool
			outputFile := f
			requireMatches := requireRE.FindAllStringSubmatch(outputFile, -1)
			for _, m := range requireMatches {
				var reverseMatch bool
				pattern, word := m[0], m[1]
				if strings.HasPrefix(word, "~") {
					word = word[1:]
					reverseMatch = true
				}
				if r, ok := requireMap[word]; ok {
					invalid = invalid || r == reverseMatch || (r != "") == reverseMatch
				}
				outputFile = strings.Replace(outputFile, pattern, "", -1)
				if invalid {
					break
				}
			}
			if invalid {
				continue
			}
			// Read file content.
			var content []byte
			if strings.HasSuffix(outputFile, ".gotmpl") {
				t := template.Must(template.New("template").Parse(templatesBox.String(f)))
				var tpl bytes.Buffer
				if err := t.Execute(&tpl, replaceMap); err != nil {
					ExitOnError(err)
				}
				content = tpl.Bytes()
			} else {
				content = templatesBox.Bytes(f)
			}
			// Create and write to new file.
			outputFile = strings.TrimSuffix(outputFile, ".gotmpl")
			if outputPath != "" {
				path, err := filepath.Abs(filepath.Join(outputPath, outputFile))
				if err != nil {
					ExitOnError(err)
				}
				outputFile = path
			}
			if folderPath := filepath.Dir(outputFile); folderPath != "." {
				if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
					ExitOnError(err)
				}
			}
			ExitOnError(ioutil.WriteFile(outputFile, content, 0644))
			fmt.Println(outputFile)
		}
	}
}
