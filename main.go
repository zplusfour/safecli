// SafeCLI, created by zplusfour.

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/pkg/browser"
)

func safe(e error) { // returns safe error
	panic(e)
}

func help() {
	fmt.Print(color.GreenString("SafeCLI help center:\n\n"))
	cmds := make(map[string]string)

	cmds[color.CyanString("get <url>")] = color.MagentaString("Gets the title of a page.")
	cmds[color.CyanString("open <url>")] = color.MagentaString("Opens a URL in your default browser.")
	cmds[color.CyanString("search <search_engine> <search_query>")] = color.MagentaString("Searchs in the search engine that you entered for the query that you entered.")

	for k := range cmds {
		fmt.Printf("%s: %s\n%s: %s\n", color.YellowString("Command"), k, color.YellowString("Description"), cmds[k])
		fmt.Println(color.HiBlackString("----------"))
	}
}

func main() {
	args := os.Args
	if len(args) < 2 {
		help()
	} else {
		switch args[1] {
		case "get":
			if len(args) < 3 {
				log.Fatal(color.RedString("Insert a URL"))
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
				fmt.Printf("%s\n", color.CyanString(data)) // sends data
			}
		case "open":
			if len(args) < 3 {
				log.Fatal(color.RedString("Insert a URL"))
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
			switch args[2] {
			case "google":
				browser.OpenURL("https://google.com/search?q=" + args[3])
			case "duckduckgo":
				browser.OpenURL("https://duckduckgo.com/?q=" + args[3])
			case "bing":
				browser.OpenURL("https://bing.com/search?q=" + args[3])
			default:
				log.Fatal(color.RedString("Not a search engine"))
			}
		default:
			log.Fatal(color.RedString("Could not find this command"))
		}
	}
}
