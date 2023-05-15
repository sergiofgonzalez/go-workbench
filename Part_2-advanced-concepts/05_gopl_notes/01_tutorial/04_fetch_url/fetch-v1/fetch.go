package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	urls := os.Args[1:]
	for _, url := range urls {
		if !strings.HasPrefix(url, `http://`) && !strings.HasPrefix(url, `https://`) {
			url = `https://` + url
		}
		fmt.Printf("Probing %s\n", url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Couldn't connect to %s: %v\n", url, err)
			os.Exit(1)
		}
		io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		fmt.Printf("\n(HTTP STATUS CODE: %d)\n", resp.StatusCode)
	}
}
