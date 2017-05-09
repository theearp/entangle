package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func fetchAndWriteActiveListings(mdb *mgo.Database) error {
	// TODO: Handle pagination from api.
	resp, err := etsyFetch(fmt.Sprintf("shops/%d/listings/active", secrets.API.ShopID))
	if err != nil {
		return fmt.Errorf("failed to fetch listings: %s", err)
	}
	defer resp.Body.Close()
	var raw *getActiveListingsResponse
	if err := json.NewDecoder(resp.Body).Decode(&raw); err != nil {
		return fmt.Errorf("failed to decode listing response: %s", err)
	}
	if len(raw.Results) == 0 {
		return errors.New("no records in etsy, this is probably wrong")
	}
	for _, l := range raw.Results {
		l.ID = bson.NewObjectId()
		if err := mdb.C("listings").Insert(l); err != nil {
			log.Printf("failed to insert listing %q: %s", l.ListingID, err)
		}
	}
	return nil
}

func (l *listing) getRemote() error {
	raw, err := etsyFetch(fmt.Sprintf("listings/%d", l.ListingID))
	if err != nil {
		return fmt.Errorf("failed to fetch listing remotely: %s", err)
	}
	defer raw.Body.Close()
	var aclr *getActiveListingsResponse
	if err := json.NewDecoder(raw.Body).Decode(&aclr); err != nil {
		return fmt.Errorf("failed to decode response: %s", err)
	}
	l = &aclr.Results[0]
	return nil
}

func (l *listing) getInventoryRemote() ([]product, error) {
	if l.ListingID == 0 {
		return nil, errors.New("Listing ID required")
	}
	raw, err := etsyFetch(fmt.Sprintf("listings/%d/inventory", l.ListingID))
	if err != nil {
		return nil, fmt.Errorf("failed to fetch inventory remotely: %s", err)
	}
	defer raw.Body.Close()
	var lir *getListingInventoryResponse
	if err := json.NewDecoder(raw.Body).Decode(&lir); err != nil {
		return nil, fmt.Errorf("failed to decode response: %s", err)
	}
	return lir.Results.Products, nil
}

func (l *listing) getImagesRemote() ([]listingImage, error) {
	if l.ListingID == 0 {
		return nil, errors.New("listing ID required")
	}
	raw, err := etsyFetch(fmt.Sprintf("listings/%d/images", l.ListingID))
	if err != nil {
		return nil, fmt.Errorf("failed to fetch inventory remotely: %s", err)
	}
	defer raw.Body.Close()
	var lir *getListingImagesResponse
	if err := json.NewDecoder(raw.Body).Decode(&lir); err != nil {
		return nil, fmt.Errorf("failed to decode response: %s", err)
	}
	return lir.Results, nil
}
