package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"gopkg.in/alecthomas/kingpin.v2"
)

// Result Dados de saida do programa
type result struct {
	StatusCode  int    `json:"code"`
	URL         string `json:"url"`
	Description string `json:"description"`
}

// Lista de urls a serem monitoradas. Implementa a interface kingpin.Settings
type urlList []url.URL

func (u *urlList) Set(value string) error {
	if url, err := url.Parse(value); err != nil {
		log.Fatalf("URL inválida: [%s] %s", value, err.Error())
	} else {
		*u = append(*u, *url)
	}
	return nil
}

func (u *urlList) String() string {
	return ""
}

func (u *urlList) IsCumulative() bool {
	return true
}

// URLList recupera a lista de URLs do comando
func URLList(s kingpin.Settings) (urls *[]url.URL) {
	urls = new([]url.URL)
	s.SetValue((*urlList)(urls))
	return
}

var (
	timeoutFlag    = kingpin.Flag("timeout", "Especifica o timeout da requisição.").Short('t').Default("10s").Duration()
	jsonFormatFlag = kingpin.Flag("json", "Saida no formato json.").Short('j').Default("false").Bool()
	urls           = URLList(kingpin.Arg("urls", "URLs para monitorar.").Required())
)

func waitHTTP(u url.URL) result {
	client := &http.Client{
		Timeout: *timeoutFlag,
	}

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		log.Printf("Problem with dial: %v.\n", err.Error())
	}

	result := result{
		StatusCode: 0,
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
	kingpin.Version("0.0.1")
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	response := waitHTTP((*urls)[0])

	if *jsonFormatFlag {
		if r, err := json.Marshal(response); err != nil {
			log.Printf("ERROR: %s", err.Error())
		} else {
			fmt.Print(string(r))
		}
	} else {
		fmt.Printf("[%s]\t %s\n", response.URL, response.Description)
	}

}
