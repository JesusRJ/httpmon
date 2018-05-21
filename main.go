package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app            = kingpin.New("httpmon", "Utilitário para monitorar disponibilidade de URLs http.")
	timeoutFlag    = app.Flag("timeout", "Especifica o timeout da requisição.").Short('t').Default("10s").Duration()
	jsonFormatFlag = app.Flag("json", "Saida no formato json.").Short('j').Default("false").Bool()
	urls           = URLList(app.Arg("urls", "URLs para monitorar.").Required())
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
	app.Version("0.0.1")
	app.Author("Reginaldo Jesus <reginaldo.jesus@gmail.com>")
	app.HelpFlag.Short('h')

	kingpin.MustParse(app.Parse(os.Args[1:]))

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
