package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

// Result Dados de saida do programa
type Result struct {
	StatusCode  int    `json:"code"`
	URL         string `json:"url"`
	Description string `json:"description"`
}

var (
	waitTimeout time.Duration
	urls        []url.URL
)

func waitHTTP(u url.URL) Result {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		log.Printf("Problem with dial: %v.\n", err.Error())
	}

	result := Result{
		StatusCode: -1,
		URL:        u.String(),
	}

	resp, err := client.Do(req)
	if err != nil {
		// log.Printf("Problem with request: %s.\n", err.Error())
		result.Description = err.Error()
	} else if err == nil && resp.StatusCode >= 200 && resp.StatusCode < 300 {
		// log.Printf("Received %d from %s\n", resp.StatusCode, u.String())
		result.StatusCode = resp.StatusCode
		result.Description = resp.Status
	}

	return result
}

func main() {
	host := "http://google.com"
	url, err := url.Parse(host)

	if err != nil {
		log.Fatalf("bad hostname provided: %s. %s", host, err.Error())
	}

	urls = append(urls, *url)

	if r, err := json.Marshal(waitHTTP(*url)); err != nil {
		log.Printf("ERROR: %s", err.Error())
	} else {
		fmt.Print(string(r))
	}

}
