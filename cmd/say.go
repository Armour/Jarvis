package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/armour/jarvis/utils"
	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
)

var sayCmd = &cobra.Command{
	Use:   "say",
	Short: "Say something（#￣▽￣#）",
	Long:  "Hi Jarvis, say something please（#￣▽￣#）",
	Run: func(cmd *cobra.Command, args []string) {
		s := spinner.New(spinner.CharSets[9], 100*time.Millisecond) // Build spinner
		s.Start()                                                   // Start the spinner
		res, err := http.Get("https://v1.hitokoto.cn/?encode=text") // Request hitokoto api
		s.Stop()                                                    // Stop the spinner
		utils.ExitOnError(err)
		defer res.Body.Close()
		contents, err := ioutil.ReadAll(res.Body)
		utils.ExitOnError(err)
		fmt.Printf("%s\n", string(contents))
	},
}

func init() {
	rootCmd.AddCommand(sayCmd)
}
