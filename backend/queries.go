package main

import (
	"fmt"
	"log"
	"time"
)

func queryProducts() ([]product, error) {
	t1 := time.Now()
	rows, err := db.Query("SELECT listing_id,	state, user_id, category_id, title, description, price, views from etsy_ActiveListings")
	if err != nil {
		return nil, fmt.Errorf("failed to get products: %s", err)
	}
	defer rows.Close()
	var products []product
	for rows.Next() {
		var p product
		if err := rows.Scan(&p.ListingID, &p.State, &p.UserID, &p.CategoryID, &p.Title, &p.Description, &p.Price, &p.Views); err != nil {
			return nil, fmt.Errorf("failed to get row: %s", err)
		}
		products = append(products, p)
	}
	log.Printf("Query took %v", time.Now().Sub(t1))
	return products, rows.Err()
}

func queryPopular() ([]product, error) {
	t1 := time.Now()
	rows, err := db.Query("SELECT listing_id,	state, user_id, category_id, title, description, price, views from etsy_ActiveListings order by views desc limit 10")
	if err != nil {
		return nil, fmt.Errorf("failed to get products: %s", err)
	}
	defer rows.Close()
	var products []product
	for rows.Next() {
		var p product
		if err := rows.Scan(&p.ListingID, &p.State, &p.UserID, &p.CategoryID, &p.Title, &p.Description, &p.Price, &p.Views); err != nil {
			return nil, fmt.Errorf("failed to get row: %s", err)
		}
		products = append(products, p)
	}
	log.Printf("Query took %v", time.Now().Sub(t1))
	return products, rows.Err()
}
