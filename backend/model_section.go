package main

import (
	"gopkg.in/mgo.v2/bson"
)

// Section represents a store section or category.
type Section struct {
	ID                 bson.ObjectId `bson:"_id" json:"id"`
	ShopSectionID      int           `json:"shop_section_id"`
	Title              string        `json:"title"`
	Rank               int           `json:"rank"`
	UserID             int           `json:"user_id"`
	ActiveListingCount int           `json:"active_listing_count"`
}

// GetSectionsResponse represents the response from esty/v2/taxonomy/categories.
type GetSectionsResponse struct {
	Count      int       `json:"count"`
	Results    []Section `json:"results"`
	Type       string    `json:"type"`
	Pagination struct {
		EffectiveLimit  int `json:"effective_limit"`
		EffectiveOffset int `json:"effective_offset"`
		NextOffset      int `json:"next_offset"`
		EffectivePage   int `json:"effective_page"`
		NextPage        int `json:"next_page"`
	} `json:"pagination"`
}
