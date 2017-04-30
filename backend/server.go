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
		msg := fmt.Sprintf("Could not get all products: %s", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
	log.Printf("fetched %d products", len(products))
	if err := json.NewEncoder(w).Encode(products); err != nil {
		http.Error(w, "failed to encode", http.StatusInternalServerError)
		return
	}
}

// PopularHandler serves the 10 most popular product by views.
func PopularHandler(w http.ResponseWriter, r *http.Request) {
	// Get a list of the most recent visits.
	products, err := queryPopular()
	if err != nil {
		msg := fmt.Sprintf("Could not get popular products: %s", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
	log.Printf("fetched %d products", len(products))
	if err := json.NewEncoder(w).Encode(products); err != nil {
		http.Error(w, "failed to encode", http.StatusInternalServerError)
		return
	}
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
	r.HandleFunc("/popular", PopularHandler)

	log.Fatal(http.ListenAndServe(":8181", handlers.CORS()(r)))
}
