package main

import (
	"encoding/json"
	"net/http"
	"time"
)

var httpClient = &http.Client{Timeout: 10 * time.Second}

// getJson is a helper function that returns the json from a URL
func getJson(url string, target interface{}) error {
	r, err := httpClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}


// getVerse returns a verse from beta.ourmanna.com API
func getVerse() (*VOTD, error) {
	votd := &VOTD{}
	err := getJson("https://beta.ourmanna.com/api/v1/get/?format=json", votd)
	if err != nil {
		return nil, err
	}
	return votd, nil
}
