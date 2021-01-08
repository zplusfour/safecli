package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	page "github.com/pkg/browser"
)

func safe(e error) {
	panic(e)
}

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatal("Please insert a URL")
	} else {
		switch args[1] {
		case "get":
			if len(args) < 3 {
				log.Fatal("Insert a URL")
			} else {
				res, err := http.Get("http://textance.herokuapp.com/title/" + args[2])
				if err != nil {
					safe(err)
				}

				body, err := ioutil.ReadAll(res.Body)
				if err != nil {
					safe(err)
				}

				data := string(body)
				// final!
				fmt.Printf("%s\n", data)
			}
		case "open":
			if len(args) < 3 {
				log.Fatal("Insert a URL")
			} else {
				if strings.HasPrefix(args[2], "http://") || strings.HasPrefix(args[2], "https://") {
					page.OpenURL(args[2])
				} else {
					err := page.OpenURL("https://" + args[2])
					if err != nil {
						safe(err)
					}
				}
			}
		}
	}
}
