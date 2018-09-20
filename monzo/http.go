package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/url"
)

func getResponse(get url.URL) io.Reader {
	req, reqErr := http.NewRequest("GET", get.String(), nil)
	if reqErr != nil {
		log.Fatal(reqErr)
	}

	req.Header.Add("User-Agent", "jlucktay (monzo-crawler)")

	res, resErr := http.DefaultClient.Do(req)
	buf := new(bytes.Buffer)

	// HTTP response errors and non-200 status codes will print to stdout and return an empty buffer
	if resErr != nil {
		errorLog.Printf("[getResponse] URL '%s' error: %v\n", get.String(), resErr)
		return buf
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		errorLog.Printf("[getResponse] URL '%s': status code error: [%d] %s\n", get.String(), res.StatusCode, res.Status)
		return buf
	}

	buf.ReadFrom(res.Body)
	return buf
}
