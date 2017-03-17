package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

const (
	defaultPaginationLimit = "100" // this goes in a url and we don't want to cast an int.
)

// URL's for the various API calls.
var (
	secrets           *config
	ShopActiveListing = "shops/%d/listings/active" // Shop ID
	ShopListing       = "shops/listing/%d"         // Listing ID
)

func urlBuilder(kind, limit, offset string, auth bool) (*url.URL, error) {
	u, err := url.Parse(secrets.API.BaseURL) // Write test for const, then remove err.
	if err != nil {
		return nil, errors.New("the url was updated, this shouldn't happen")
	}

	q := u.Query()
	if auth {
		q.Set("api_key", secrets.API.Key)
	}

	if limit != "" {
		q.Set("limit", limit)
		q.Set("offset", offset)
	}

	switch kind {
	case "GetActiveListings":
		u.Path = fmt.Sprintf("%s/%s", u.Path, fmt.Sprintf(ShopActiveListing, secrets.API.ShopID))
	default:
		return nil, fmt.Errorf("kind %q does not match available methods", kind)
	}
	u.RawQuery = q.Encode()
	return u, nil
}

func etsyFetch(req *http.Request) (*http.Response, error) {
	client := &http.Client{}
	req.Header.Add("Cookie", secrets.API.UserPrefs) // Adjust if necessary.
	if err := req.ParseForm(); err != nil {
		return nil, err
	}
	fmt.Printf("the url being fetched: %s\n", req.URL)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetActiveListings retrieves all active listings.
func GetActiveListings() (*GetActiveListingResponse, error) {
	loc, err := urlBuilder("GetActiveListings", defaultPaginationLimit, "0", true)
	if err != nil {
		return nil, fmt.Errorf("failed to build url: %s", err)
	}
	fmt.Printf("the path going in, %s\n", loc.Path)
	req, err := http.NewRequest("GET", loc.String(), nil)
	if err != nil {
		return nil, err
	}
	resp, err := etsyFetch(req)
	if err != nil {
		return nil, fmt.Errorf("failed to etsyFetch: %s", err)
	}
	defer resp.Body.Close()
	var result *GetActiveListingResponse
	// decode
	// respBody, _ := ioutil.ReadAll(resp.Body)
	// fmt.Print(string(respBody))
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %s", err)
	}

	// paginate
	// if result.Count > 100 {
	// 	var wg sync.WaitGroup
	// 	counter := result.Count
	// 	for counter > 0 {
	// 		wg.Add(1)
	// 		go func(c int) {
	// 			var tmp *GetActiveListingResponse
	// 			turl, err := urlBuilder("GetActiveListings", defaultPaginationLimit, c, true)
	// 			if err == nil {
	// 				tr, err := etsyFetch(turl)
	// 			}

	// 			if err != nil {
	// 				fmt.Printf("failed to fetch pagninated urls: %s", err)
	// 			}
	// 		}(strconv.Itoa(counter))
	// 		counter = counter - defaultPaginationLimit
	// 	}
	// 	wg.Done()
	// }

	return result, nil
}

func main() {
	var err error
	if secrets, err = getConfig("config.yaml"); err != nil {
		log.Fatalf("failed to collect secrets: %s", err)
	}
	listings, err := GetActiveListings()
	if err != nil {
		fmt.Printf("failed to get active listings: %s\n", err)
	} else {
		fmt.Printf("the listings: %#v\n", listings)
	}
}
