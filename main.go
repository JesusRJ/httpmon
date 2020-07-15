package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
	"net/http"
	"net/url"
	"os"
)

// result Result output
type result struct {
	StatusCode        int    `json:"code"`
	StatusDescription string `json:"description"`
	URL               string `json:"url"`
}

var (
	app               = kingpin.New("httpmon", "Utilitário para monitorar disponibilidade de URLs http.")
	timeoutFlag       = app.Flag("timeout", "Especifica o timeout da requisição.").Short('t').Default("5s").Duration()
	verboseFormatFlag = app.Flag("verbose", "Imprime mais informações.").Short('v').Default("false").Bool()
	jsonFormatFlag    = app.Flag("json", "Saida no formato json.").Short('j').Default("false").Bool()
	urlFlag           = app.Flag("url", "URL a monitorar.").Short('u').Required().String()
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
		result.StatusDescription = err.Error()
	} else {
		result.StatusCode = resp.StatusCode
		result.StatusDescription = resp.Status
	}

	return result
}

func main() {
	app.Version("0.0.1")
	app.Author("Reginaldo Jesus <reginaldo.jesus@gmail.com>")
	app.HelpFlag.Short('h')

	kingpin.MustParse(app.Parse(os.Args[1:]))

	u, err := url.Parse(*urlFlag)

	if err != nil {
		log.Fatalf("Invalid URL [%s] %s", *urlFlag, err.Error())
	}

	response := waitHTTP(*u)

	switch {
	case *jsonFormatFlag:
		if r, err := json.Marshal(response); err != nil {
			log.Printf("ERROR: %s\n", err.Error())
		} else {
			fmt.Println(string(r))
		}
	case *verboseFormatFlag:
		log.Printf("[%s]\t %s\n", response.URL, response.StatusDescription)
	default:
		switch {
		case response.StatusCode >= 200 && response.StatusCode < 300:
			fmt.Println("URL disponível")
		default:
			fmt.Println("URL indisponível")
		}
	}

}
