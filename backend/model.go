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

func (p *product) getProduct(db *sql.DB) error {
	return errors.New("Not Implmeneted")
}

func (p *product) updateProduct(db *sql.DB) error {
	return errors.New("Not Implmeneted")
}

func (p *product) deleteProduct(db *sql.DB) error {
	return errors.New("Not Implmeneted")
}

func (p *product) createProduct(db *sql.DB) error {
	return errors.New("Not Implmeneted")
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
