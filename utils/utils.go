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

	"github.com/gobuffalo/packr"
)

const (
	requirePattern = `\[(.*)\]`
)

var (
	requireRE = regexp.MustCompile(requirePattern)
)

// ExitOnError prints the error message and exit with code 1
func ExitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// RunCommand executes the input command
func RunCommand(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	var outBuf bytes.Buffer
	var errBuf bytes.Buffer
	cmd.Stdout = &outBuf
	cmd.Stderr = &errBuf
	if err := cmd.Run(); err != nil {
		if errBuf.String() != "" {
			fmt.Printf("%q", errBuf.String())
		}
		ExitOnError(err)
	}
	if outBuf.String() != "" {
		fmt.Printf("%q", outBuf.String())
	}
}

// GenerateFile generates files from template folder
func GenerateFile(templatePath string, outputPath string, requireMap map[string]interface{}, replaceMap map[string]interface{}) {
	fmt.Println("Generated file:")
	templatesBox := packr.NewBox(templatePath)
	for _, f := range templatesBox.List() {
		var missRequired bool
		outputFile := f

		// Check requirement
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

		// Get file content
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

		// Create and write to new file
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
