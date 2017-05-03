package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
)

func queryProducts(db *sql.DB, catID string) ([]product, error) {
	t1 := time.Now()
	stmt := "SELECT listing_id,	state, user_id, category_id, title, description, price, views from etsy_ActiveListings"
	if catID != "" {
		stmt = fmt.Sprintf("SELECT listing_id,	state, user_id, category_id, title, description, price, views from etsy_ActiveListings where category_id = %s", catID)
	}
	rows, err := db.Query(stmt)
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

func queryPopular(db *sql.DB) ([]product, error) {
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

func queryProduct(db *sql.DB, id int) (product, error) {
	t1 := time.Now()
	var p product
	rows, err := db.Query(fmt.Sprintf("SELECT listing_id,	state, user_id, category_id, title, description, price, views from etsy_ActiveListings where listing_id = %d", id))
	if err != nil {
		return p, fmt.Errorf("failed to get products: %s", err)
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&p.ListingID, &p.State, &p.UserID, &p.CategoryID, &p.Title, &p.Description, &p.Price, &p.Views); err != nil {
			return p, fmt.Errorf("failed to get row: %s", err)
		}
	}
	log.Printf("Query took %v", time.Now().Sub(t1))
	return p, rows.Err()
}

func (p *product) getProduct(db *sql.DB) error {
	t1 := time.Now()
	rows, err := db.Query(fmt.Sprintf("SELECT listing_id,	state, user_id, category_id, title, description, price, views from etsy_ActiveListings where listing_id = %d", p.ListingID))
	if err != nil {
		return fmt.Errorf("failed to get products: %s", err)
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&p.ListingID, &p.State, &p.UserID, &p.CategoryID, &p.Title, &p.Description, &p.Price, &p.Views); err != nil {
			return fmt.Errorf("failed to get row: %s", err)
		}
	}
	log.Printf("Query took %v", time.Now().Sub(t1))
	return rows.Err()
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

func queryCategories(db *sql.DB) ([]category, error) {
	t1 := time.Now()
	rows, err := db.Query(`SELECT 
			category_id,			
			name,				    
			meta_title,			
			meta_keywords,		
			meta_description,
			page_description,
			page_title,			
			category_name,		
			short_name,			
			long_name,				
			num_children from etsy_Categories`)
	if err != nil {
		return nil, fmt.Errorf("failed to get categories: %s", err)
	}
	defer rows.Close()
	var categories []category
	for rows.Next() {
		var c category
		if err := rows.Scan(
			&c.CategoryID,
			&c.Name,
			&c.MetaTitle,
			&c.MetaKeywords,
			&c.MetaDescription,
			&c.PageDescription,
			&c.PageTitle,
			&c.CategoryName,
			&c.ShortName,
			&c.LongName,
			&c.NumChildren,
		); err != nil {
			return nil, fmt.Errorf("failed to get row: %s", err)
		}
		categories = append(categories, c)
	}
	log.Printf("Query took %v", time.Now().Sub(t1))
	return categories, rows.Err()
}

func querySections(db *sql.DB) ([]section, error) {
	t1 := time.Now()
	rows, err := db.Query(`SELECT 
			shop_section_id,			
			title,	
			rank,								
			user_id,							
			active_listing_count from etsy_Sections`)
	if err != nil {
		return nil, fmt.Errorf("failed to get sections: %s", err)
	}
	defer rows.Close()
	var sections []section
	for rows.Next() {
		var s section
		if err := rows.Scan(
			&s.ShopSectionID,
			&s.Title,
			&s.Rank,
			&s.UserID,
			&s.ActiveListingCount,
		); err != nil {
			return nil, fmt.Errorf("failed to get row: %s", err)
		}
		sections = append(sections, s)
	}
	log.Printf("Query took %v", time.Now().Sub(t1))
	return sections, rows.Err()
}

func (c *category) getCategory(db *sql.DB) error {
	return errors.New("Not Implmeneted")
}

func (c *category) updateCategory(db *sql.DB) error {
	return errors.New("Not Implmeneted")
}

func (c *category) deleteCategory(db *sql.DB) error {
	return errors.New("Not Implmeneted")
}

func (c *category) createCategory(db *sql.DB) error {
	return errors.New("Not Implmeneted")
}
