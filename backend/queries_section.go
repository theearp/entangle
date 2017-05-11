package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func syncSections(mdb *mgo.Database) error {
	// TODO: Handle pagination from api.
	resp, err := etsyFetch(fmt.Sprintf("/shops/%d/sections", secrets.API.ShopID))
	if err != nil {
		return fmt.Errorf("failed to fetch listings: %s", err)
	}
	defer resp.Body.Close()
	var raw *GetSectionsResponse
	if err := json.NewDecoder(resp.Body).Decode(&raw); err != nil {
		return fmt.Errorf("failed to decode listing response: %s", err)
	}
	if len(raw.Results) == 0 {
		return errors.New("no records in etsy, this is probably wrong")
	}
	for _, l := range raw.Results {
		l.ID = bson.NewObjectId()
		if err := mdb.C("sections").Insert(l); err != nil {
			log.Printf("failed to insert listing %q: %s", l.ShopSectionID, err)
		}
	}
	return nil
}
