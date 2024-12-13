package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func DownloadInput(url string) string {
	// Function to download input.
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	content, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}

func DownloadInputWithCookie(url string, sessionCookie string) string {
	// Initializing the HTTP Client.
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	// Adding cookie as per curl request.
	cookieCurlString := fmt.Sprintf("session=%s", sessionCookie)
	req.Header.Add("Cookie", cookieCurlString)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	// Parsing the download
	content, err := io.ReadAll(resp.Body)

	return string(content)
}
