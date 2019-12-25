package main

import (
	"io/ioutil"
	"net/http"
	"strings"
	//"time"
)

// User supplied constants
var (
	subdomain  string
	domainName string
	ddnsToken  string
)

// Execution constants
var urls []string = []string{
	"https://api.ipify.org",
	"https://icanhazip.com",
	"https://bot.whatismyipaddress.com",
}
var ddnsURL string = "https://dynamicdns.park-your-domain.com/update"
var clientIP, updateURL string

func getClientIP(urls []string) string {
	var resp *http.Response
	var err error
	for _, v := range urls {
		resp, err = http.Get(v)
		if err == nil {
			body, _ := ioutil.ReadAll(resp.Body)
			return string(body)
		}
		continue
	}
	panic(err)
}

func main() {
	clientIP := getClientIP(urls)
	updateURL = strings.Join([]string{
		ddnsURL,
		"?host=", subdomain,
		"&domain=", domainName,
		"&password=", ddnsToken,
		"&ip=", clientIP,
	}, "")

	_, err := http.Get(updateURL)
	if err != nil {
		panic(err)
	}
}
