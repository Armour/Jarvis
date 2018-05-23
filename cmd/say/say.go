package say

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	c "github.com/armour/jarvis/cmd"
	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
)

var sayCmd = &cobra.Command{
	Use:   "say",
	Short: "Say something.（#￣▽￣#）",
	Long:  "Hi Jarvis, say something please.（#￣▽￣#）",
	Run: func(cmd *cobra.Command, args []string) {
		s := spinner.New(spinner.CharSets[9], 100*time.Millisecond) // Build spinner
		s.Start()                                                   // Start the spinner
		res, err := http.Get("https://v1.hitokoto.cn/?encode=text") // Request hitokoto api
		s.Stop()                                                    // Stop the spinner
		c.ExitOnError(err)
		defer res.Body.Close()
		contents, err := ioutil.ReadAll(res.Body)
		c.ExitOnError(err)
		fmt.Printf("%s\n", string(contents))
	},
}

func init() {
	c.RootCmd.AddCommand(sayCmd)
}
