package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	urls := os.Args[1:]
	for _, url := range urls {
		

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Couldn't connect to %s: %v\n", url, err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Couldn't exhaust responsse body from %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", b)
	}
}