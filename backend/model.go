package main

import (
	"database/sql"
	"errors"
)

type product struct {
	ListingID   int    `json:"listing_id"`
	State       string `json:"state"`
	UserID      int    `json:"user_id"`
	CategoryID  int    `json:"category_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       string `json:"price"`
	Views       int64  `json:"views"`
}

type category struct {
	CategoryID      int    `json:"category_id"`
	Name            string `json:"name"`
	MetaTitle       string `json:"meta_title"`
	MetaKeywords    string `json:"meta_keywords"`
	MetaDescription string `json:"meta_description"`
	PageDescription string `json:"page_description"`
	PageTitle       string `json:"page_title"`
	CategoryName    string `json:"category_name"`
	ShortName       string `json:"short_name"`
	LongName        string `json:"long_name"`
	NumChildren     int    `json:"num_children"`
}

type cart struct {
}

func (c *cart) getCart(db *sql.DB) error {
	return errors.New("Not Implmeneted")
}

func (c *cart) updateCart(db *sql.DB) error {
	return errors.New("Not Implmeneted")
}

func (c *cart) deleteCart(db *sql.DB) error {
	return errors.New("Not Implmeneted")
}

func (c *cart) createCart(db *sql.DB) error {
	return errors.New("Not Implmeneted")
}
