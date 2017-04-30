package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var (
	secrets *config
	db      *sql.DB
	dbName  = "playground"
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

// ProductHandler serves a single product given its listing id.
func ProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sid := vars["id"]
	id, err := strconv.Atoi(sid)
	if err != nil {
		http.Error(w, "failed to convert listing id string to int", http.StatusInternalServerError)
		return
	}
	product, err := queryProduct(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to retrieve product: %d", id), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(product); err != nil {
		http.Error(w, "failed to encode", http.StatusInternalServerError)
		return
	}
}

func main() {
	var err error
	if secrets, err = getSQLConfig("config.yaml"); err != nil {
		log.Fatalf("failed to get secrets: %s", err)
	}
	if err = connect("local"); err != nil {
		log.Fatalf("failed to connect: %s", err)
	}
	log.Println("db connected successfully.")
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/products", ProductsHandler)
	r.HandleFunc("/popular", PopularHandler)
	r.HandleFunc("/product/{id}", ProductHandler)

	log.Fatal(http.ListenAndServe(":8181", handlers.CORS()(r)))
}
