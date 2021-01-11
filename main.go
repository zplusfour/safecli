// SafeCLI, created by zplusfour.

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/pkg/browser"
)

func safe(e error) { // returns safe error
	panic(e)
}

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatal("Please insert a command")
	} else {
		switch args[1] {
		case "get":
			if len(args) < 3 {
				log.Fatal("Insert a URL")
			} else {
				res, err := http.Get("http://textance.herokuapp.com/title/" + args[2]) // GETs the website title from Textance
				if err != nil {
					safe(err)
				}

				body, err := ioutil.ReadAll(res.Body)
				if err != nil {
					safe(err)
				}

				data := string(body)
				fmt.Printf("%s\n", data) // sends data
			}
		case "open":
			if len(args) < 3 {
				log.Fatal("Insert a URL")
			} else {
				if strings.HasPrefix(args[2], "http://") || strings.HasPrefix(args[2], "https://") { // scans if it has https:// or http:// then redirect
					browser.OpenURL(args[2])
				} else { // else, open it on https://
					err := browser.OpenURL("https://" + args[2])
					if err != nil {
						safe(err)
					}
				}
			}
        case "search":
            switch args[2]{
                case "google":
                    browser.OpenURL("https://google.com/search?q="+args[3])
                case "duckduckgo":
                    browser.OpenURL("https://duckduckgo.com/?q="+args[3])
                case "bing":
                    browser.OpenURL("https://bing.com/search?q="+args[3])
                default:
                    log.Fatal("Not a search engine")
            }
        default:
            log.Fatal("Could not find this command")
		}
	}
}
