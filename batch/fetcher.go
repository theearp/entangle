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

var (
	secrets           *config
	shopActiveListing = "shops/%d/listings/active" // Shop ID
	shopListing       = "shops/listing/%d"         // Listing ID
	category          = "taxonomy/categories"
	section           = "/shops/%d/sections"
	images            = "/listings/%d/images"
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
		u.Path = fmt.Sprintf("%s/%s", u.Path, fmt.Sprintf(shopActiveListing, secrets.API.ShopID))
	case "GetCategories":
		u.Path = fmt.Sprintf("%s/%s", u.Path, category)
	case "GetSections":
		u.Path = fmt.Sprintf("%s/%s", u.Path, fmt.Sprintf(section, secrets.API.ShopID))
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
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %s", err)
	}
	return result, nil
}

// GetCategories get the top level categories from etsy.
func GetCategories() (*GetCategoriesResponse, error) {
	loc, err := urlBuilder("GetCategories", "", "0", true)
	if err != nil {
		return nil, fmt.Errorf("failed to build url: %s", err)
	}
	req, err := http.NewRequest("GET", loc.String(), nil)
	if err != nil {
		return nil, err
	}
	resp, err := etsyFetch(req)
	if err != nil {
		return nil, fmt.Errorf("failed to etsyFetch: %s", err)
	}
	defer resp.Body.Close()
	var result *GetCategoriesResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %s", err)
	}
	return result, nil
}

// GetSections get the shop sections from etsy, sections are what I thought categories were.
func GetSections() (*GetSectionsResponse, error) {
	loc, err := urlBuilder("GetSections", "", "0", true)
	if err != nil {
		return nil, fmt.Errorf("failed to build url: %s", err)
	}
	req, err := http.NewRequest("GET", loc.String(), nil)
	if err != nil {
		return nil, err
	}
	resp, err := etsyFetch(req)
	if err != nil {
		return nil, fmt.Errorf("failed to etsyFetch: %s", err)
	}
	defer resp.Body.Close()
	var result *GetSectionsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %s", err)
	}
	return result, nil
}

func main() {
	var err error
	if secrets, err = getConfig("config.yaml"); err != nil {
		log.Fatalf("failed to collect secrets: %s", err)
	}
	if err = connect("local"); err != nil {
		log.Fatalf("failed to connect to database: %s", err)
	}
	// if err = createListingTable(); err != nil {
	// 	log.Fatalf("failed to create table: %s", err)
	// }
	// if err = createCategoryTable(); err != nil {
	// 	log.Fatalf("failed to create table: %s", err)
	// }
	if err = createSectionsTable(); err != nil {
		log.Fatalf("failed to create table: %s", err)
	}
	log.Println("successfully created table")
	// listings, err := GetActiveListings()
	// if err != nil {
	// 	fmt.Printf("failed to get active listings: %s\n", err)
	// } else {
	// 	fmt.Printf("the listings: %#v\n", listings)
	// }
	// if err = writeListings(listings); err != nil {
	// 	log.Fatalf("failed to write listings to database: %s", err)
	// }
	// categories, err := GetCategories()
	// if err != nil {
	// 	fmt.Printf("failed to get categories: %s\n", err)
	// } else {
	// 	fmt.Printf("the categories: %#v\n", categories)
	// }
	// if err = writeCategories(categories); err != nil {
	// 	log.Fatalf("failed to write categories to database: %s", err)
	// }
	sections, err := GetSections()
	if err != nil {
		fmt.Printf("failed to get sections: %s\n", err)
	} else {
		fmt.Printf("the sections: %#v\n", sections)
	}
	if err = writeSections(sections); err != nil {
		log.Fatalf("failed to write sections to database: %s", err)
	}
}
