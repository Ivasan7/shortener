/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"runtime"
	"os/exec"
	"log"
	"github.com/spf13/cobra"
	"net/http"
	"io/ioutil"
)

var openInBrowser bool
// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Fetching long or short URLs",
	Long: `Fetching URLs either from conventional format or short format.
	 The short format first has to be generated. For example:

	./shortener get -l google.com OR
	./shortener get -s fake.it/1
	
	Running the app natively gives the option to open the URL in the browser
	by providing the -o flag
	./shortener get -l google.com -o`,
	Run: func(cmd *cobra.Command, args []string) {
		shortUrlName := cmd.Flag("shortUrl")
		longUrlName := cmd.Flag("longUrl")
		if shortUrlName.Value.String() == "" && longUrlName.Value.String() == "" {
			fmt.Println("Short or Long Url name must be provided")
		} else {
			if longUrlName.Value.String() != "" {
				// TODO remove duplication of code
				if openInBrowser {
					openbrowser(longUrlName.Value.String())
				} else {
					readAsText("http://" +longUrlName.Value.String())
				}
			} else {
				ID,longUrl := DB.GetLongUrl(shortUrlName.Value.String())
				if ID == -1 {
					log.Fatalf("URL is not recognised: %s", longUrl)
				}
				if openInBrowser {
					openbrowser(longUrl)
				} else {
					log.Printf("Fetching %s",longUrl)
					readAsText("http://" + longUrl)	
				}
			}
		}	
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().StringP("shortUrl","s","", "short URL name")
	getCmd.Flags().StringP("longUrl","l","", " long URL name")
	getCmd.Flags().BoolVar(&openInBrowser,"o",false,"open in browser")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}

func readAsText(url string) {
	fmt.Printf("HTML code of %s ...\n", url)
	resp, err := http.Get(url)
	// handle the error if there is one
	if err != nil {
		panic(err)
	}
	// do this now so it won't be forgotten
	defer resp.Body.Close()
	// reads html as a slice of bytes
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	// show the HTML code as a string %s
	fmt.Printf("%s\n", html)

}