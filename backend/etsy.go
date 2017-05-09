package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func etsyFetch(resource string) (*http.Response, error) {
	u, err := url.Parse(secrets.API.BaseURL)
	if err != nil {
		return nil, errors.New("the url was updated, this shouldn't happen")
	}

	q := u.Query()
	q.Set("api_key", secrets.API.Key)

	u.Path = fmt.Sprintf("%s/%s", u.Path, resource)
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	req.Header.Add("Cookie", secrets.API.UserPrefs) // Adjust if necessary.
	if err := req.ParseForm(); err != nil {
		return nil, err
	}
	log.Printf("the url being fetched: %s\n", req.URL)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
