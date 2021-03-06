package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

func connect(env string) error {
	cfg, err := secrets.env(env)
	if err != nil {
		return fmt.Errorf("failed to collect secrets: %s", err)
	}

	if env == "local" {
		if db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", cfg.User, cfg.Passwd, cfg.Addr, cfg.DBName)); err != nil {
			return fmt.Errorf("failed to connect to db: %s", err)
		}
	} else {
		if db, err = mysql.DialCfg(cfg); err != nil {
			return fmt.Errorf("failed to connect to db: %s", err)
		}
	}

	return db.Ping()
}

func createSectionsTable() error {
	stmt := `CREATE TABLE IF NOT EXISTS etsy_Sections (
							shop_section_id				BIGINT,
							title									VARCHAR(255),
							rank									BIGINT,
							user_id								BIGINT,
							active_listing_count 	BIGINT
	)`
	_, err := db.Exec(stmt)
	return err
}

func createCategoryTable() error {
	stmt := `CREATE TABLE IF NOT EXISTS etsy_Categories (
							category_id				BIGINT,		
							name				    	VARCHAR(255),
							meta_title				VARCHAR(255),
							meta_keywords			VARCHAR(255),
							meta_description	VARCHAR(255),
							page_description	VARCHAR(255),
							page_title				VARCHAR(255),
							category_name			VARCHAR(255),
							short_name				VARCHAR(255),
							long_name					VARCHAR(255),
							num_children			BIGINT	
		)`
	_, err := db.Exec(stmt)
	return err
}

func createListingTable() error {
	stmt := `CREATE TABLE IF NOT EXISTS etsy_ActiveListings (
          		listing_id						BIGINT,
              state     						VARCHAR(255),
							user_id 							BIGINT,
							category_id 					BIGINT,
							title 								VARCHAR(255),
							description 					VARCHAR(8000),
							creation_tsz 					BIGINT,
							ending_tsz 						BIGINT,
							original_creation_tsz BIGINT,
							price 								VARCHAR(255),
							currency_code 				VARCHAR(255),
							quantity 							BIGINT,		
							shop_section_id				BIGINT,
							state_tsz 						BIGINT,
							url										VARCHAR(255),
							views								  BIGINT,
							num_favors						BIGINT,
							processing_min 				BIGINT,
							processing_max 				BIGINT,
							who_made 							VARCHAR(255),
							is_supply 						VARCHAR(255),
							when_made 						VARCHAR(255),
							item_dimensions_unit 	VARCHAR(255),
							is_private 						BOOLEAN,
							file_data 						VARCHAR(255),
							language 							VARCHAR(255),
							has_variations 				BOOLEAN,
							taxonomy_id 					BIGINT
          )`
	_, err := db.Exec(stmt)
	return err
}

func writeListings(l *GetActiveListingResponse) error {
	stmt := `INSERT INTO etsy_ActiveListings (
			listing_id,						
			state,     						
			user_id, 							
			category_id, 					
			title, 								
			description,
			price,
			views) VALUES (?,?,?,?,?,?,?,?)`
	for _, listing := range l.Results {
		_, err := db.Exec(stmt,
			listing.ListingID,
			listing.State,
			listing.UserID,
			listing.CategoryID,
			listing.Title,
			listing.Description,
			listing.Price,
			listing.Views)
		if err != nil {
			log.Printf("failed to write stmt: %s", stmt)
		}
	}
	return nil
}

func writeCategories(c *GetCategoriesResponse) error {
	stmt := `INSERT INTO etsy_Categories (
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
			num_children) VALUES (?,?,?,?,?,?,?,?,?,?,?)`
	for _, category := range c.Results {
		_, err := db.Exec(stmt,
			category.CategoryID,
			category.Name,
			category.MetaTitle,
			category.MetaKeywords,
			category.MetaDescription,
			category.PageDescription,
			category.PageTitle,
			category.CategoryName,
			category.ShortName,
			category.LongName,
			category.NumChildren)
		if err != nil {
			log.Printf("failed to write stmt: %s", stmt)
		}
	}
	return nil
}

func writeSections(c *GetSectionsResponse) error {
	stmt := `INSERT INTO etsy_Sections (
			shop_section_id,			
			title,	
			rank,								
			user_id,							
			active_listing_count) VALUES (?,?,?,?,?)`
	for _, section := range c.Results {
		_, err := db.Exec(stmt,
			section.ShopSectionID,
			section.Title,
			section.Rank,
			section.UserID,
			section.ActiveListingCount)
		if err != nil {
			log.Printf("failed to write stmt: %s", stmt)
		}
	}
	return nil
}
