package main

import (
	"log"
	"net/url"

	"gopkg.in/alecthomas/kingpin.v2"
)

// result Dados de saida do programa.
type result struct {
	StatusCode  int    `json:"code"`
	URL         string `json:"url"`
	Description string `json:"description"`
}

// urlList Lista de urls a serem monitoradas. Implementa a interface kingpin.Settings.
type urlList []url.URL

func (u *urlList) Set(value string) error {
	if url, err := url.Parse(value); err != nil {
		log.Fatalf("Invalid URL [%s] %s", value, err.Error())
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
