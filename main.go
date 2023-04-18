package main

import (
	"fmt"
	"net/http"
	"strings"
)

func fetchHeaders(url string, modifyHeadersFunc func(req *http.Request)) (http.Header, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	modifyHeadersFunc(req)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return resp.Header, nil
}

func modifyHeaders(req *http.Request) {
	// Add or modify headers here, for example:
	req.Header.Set("User-Agent", "Custom-User-Agent")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Add("Custom-Header", "Custom-Value")
}

func main() {
	url := "https://example.com"
	headers, err := fetchHeaders(url, modifyHeaders)
	if err != nil {
		fmt.Println("Error fetching headers:", err)
		return
	}

	for key, values := range headers {
		fmt.Printf("%s: %s\n", key, strings.Join(values, ", "))
	}
}
