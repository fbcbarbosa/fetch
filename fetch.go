// Fetch prints the content found at a URL.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		err := Fetch(url, os.Stdout)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fetch: %v\n", err)
			os.Exit(1)
		}
	}
}

// Fetch prints the content found at a URL to the given Writer.
func Fetch(url string, dst io.Writer) error {
	if !strings.HasPrefix(url, "https://") && !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}

	resp, err := http.Get(url)

	if err != nil {
		return err
	}

	_, err = io.Copy(dst, resp.Body)
	resp.Body.Close()

	return err
}
