package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"time"

	gmysql "github.com/go-sql-driver/mysql"
)

var (
	secrets *config
	db      *sql.DB
	dbName  = "playground"
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

// HomeHandler serves a welcome.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

// ProductsHandler serves all products from the database.
func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	// Get a list of the most recent visits.
	products, err := queryProducts()
	if err != nil {
		msg := fmt.Sprintf("Could not get products: %s", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
	log.Printf("fetched %d products", len(products))
	if err := json.NewEncoder(w).Encode(products); err != nil {
		http.Error(w, "failed to encode", http.StatusInternalServerError)
		return
	}
}

type product struct {
	ListingID   int    `json:"listing_id"`
	State       string `json:"state"`
	UserID      int    `json:"user_id"`
	CategoryID  int    `json:"category_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func queryProducts() ([]product, error) {
	t1 := time.Now()
	rows, err := db.Query("SELECT listing_id,	state, user_id, category_id, title, description from etsy_ActiveListings")
	if err != nil {
		return nil, fmt.Errorf("failed to get products: %s", err)
	}
	defer rows.Close()
	var products []product
	for rows.Next() {
		var p product
		if err := rows.Scan(&p.ListingID, &p.State, &p.UserID, &p.CategoryID, &p.Title, &p.Description); err != nil {
			return nil, fmt.Errorf("failed to get row: %s", err)
		}
		products = append(products, p)
	}
	log.Printf("Query took %v", time.Now().Sub(t1))
	return products, rows.Err()
}

func main() {
	var err error
	if secrets, err = getConfig("config.yaml"); err != nil {
		log.Fatalf("failed to get secrets: %s", err)
	}
	if err = connect(); err != nil {
		log.Fatalf("failed to connect: %s", err)
	}
	log.Println("db connected successfully.")
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/products", ProductsHandler)

	log.Fatal(http.ListenAndServe(":8181", handlers.CORS()(r)))
}
