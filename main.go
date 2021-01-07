package main

import (
	"net/http"
	"os"
	"log"
	"io/ioutil"
	"fmt"
)
func safe(e error) {
	panic(e)
}

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatal("Please insert a URL")
	} else if len(args) > 2 {
		log.Fatal("Too many arguments")
	} else {
		res, err := http.Get("http://textance.herokuapp.com/title/"+args[1])
		if err != nil {
			safe(err)
		}

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			safe(err)
		}

		data := string(body)
		// final!
		fmt.Println(data)
	}
}