package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"

	gmysql "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

func connect() error {
	var err error
	cfg := &gmysql.Config{
		Addr:   secrets.SQL.Address,
		User:   secrets.SQL.Username,
		Passwd: secrets.SQL.Password,
		DBName: secrets.SQL.DBName,
	}
	if db, err = mysql.DialCfg(cfg); err != nil {
		return fmt.Errorf("failed to connect to db: %s", err)
	}
	return db.Ping()
}

func createTable() error {
	stmt := `CREATE TABLE etsy_ActiveListings (
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
