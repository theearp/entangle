package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Entangle represents the Entangle API service.
type Entangle struct {
	Config string
	DB     *sql.DB
	Router *mux.Router
}

func (e *Entangle) startUp(env string) error {
	e.resgisterRoutes()
	secrets, err := getSQLConfig(e.Config)
	if err != nil {
		return fmt.Errorf("failed to get secrets: %s", err)
	}
	cfg, err := secrets.env(env)
	if err != nil {
		return fmt.Errorf("failed to collect secrets: %s", err)
	}

	log.Println("connecting to database....")
	if env == "local" {
		if e.DB, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", cfg.User, cfg.Passwd, cfg.Addr, cfg.DBName)); err != nil {
			return fmt.Errorf("failed to connect to db: %s", err)
		}
	} else {
		if e.DB, err = mysql.DialCfg(cfg); err != nil {
			return fmt.Errorf("failed to connect to db: %s", err)
		}
	}
	return e.DB.Ping()
}

func (e *Entangle) resgisterRoutes() {
	e.Router = mux.NewRouter()
	e.Router.HandleFunc("/", e.home).Methods("GET")
	e.Router.HandleFunc("/products", e.products).Methods("GET")
	e.Router.HandleFunc("/popular", e.popularProducts).Methods("GET")
	e.Router.HandleFunc("/product/{id}", e.product).Methods("GET")
	e.Router.HandleFunc("/categories", e.categories).Methods("GET")
	e.Router.HandleFunc("/product_category/{id}", e.productCategory).Methods("GET")
	e.Router.HandleFunc("/sections", e.sections).Methods("GET")
}

func (e *Entangle) run(addr string) {
	log.Println("Serving API")
	log.Fatal(http.ListenAndServe(addr, handlers.CORS()(e.Router)))
}

func (e *Entangle) home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Nothing here")
}

func (e *Entangle) products(w http.ResponseWriter, r *http.Request) {
	t1 := time.Now()
	products, err := queryProducts(e.DB, "")
	if err != nil {
		msg := fmt.Sprintf("Could not get all products: %s", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
	if err := renderJSON(w, products); err != nil {
		http.Error(w, "failed to encode", http.StatusInternalServerError)
		return
	}
	log.Printf("%d /products returned successfully in %v", len(products), time.Now().Sub(t1))
}

func (e *Entangle) popularProducts(w http.ResponseWriter, r *http.Request) {
	t1 := time.Now()
	// Get a list of the most recent visits.
	products, err := queryPopular(e.DB)
	if err != nil {
		msg := fmt.Sprintf("Could not get popular products: %s", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
	if err := renderJSON(w, products); err != nil {
		http.Error(w, "failed to encode", http.StatusInternalServerError)
		return
	}
	log.Printf("%d /popular returned successfully in %v", len(products), time.Now().Sub(t1))
}

func (e *Entangle) product(w http.ResponseWriter, r *http.Request) {
	t1 := time.Now()
	vars := mux.Vars(r)
	sid := vars["id"]
	id, err := strconv.Atoi(sid)
	if err != nil {
		http.Error(w, "failed to convert listing id string to int", http.StatusInternalServerError)
		return
	}
	product, err := queryProduct(e.DB, id)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to retrieve product: %d", id), http.StatusInternalServerError)
		return
	}
	if err := renderJSON(w, product); err != nil {
		http.Error(w, "failed to encode", http.StatusInternalServerError)
		return
	}
	log.Printf("/product/%d returned successfully in %v", id, time.Now().Sub(t1))
}

func (e *Entangle) categories(w http.ResponseWriter, r *http.Request) {
	t1 := time.Now()
	cs, err := queryCategories(e.DB)
	if err != nil {
		http.Error(w, fmt.Sprintf("could not get categories: %s", err), http.StatusInternalServerError)
		return
	}
	if err := renderJSON(w, cs); err != nil {
		http.Error(w, "failed to encode", http.StatusInternalServerError)
		return
	}
	log.Printf("%d /categories returned successfully in %v", len(cs), time.Now().Sub(t1))
}

func (e *Entangle) productCategory(w http.ResponseWriter, r *http.Request) {
	t1 := time.Now()
	vars := mux.Vars(r)
	products, err := queryProducts(e.DB, vars["id"])
	if err != nil {
		http.Error(w, fmt.Sprintf("could not get products: %s", err), http.StatusInternalServerError)
		return
	}
	if err := renderJSON(w, products); err != nil {
		http.Error(w, "failed to encode", http.StatusInternalServerError)
		return
	}
	log.Printf("%d /products returned successfully in %v", len(products), time.Now().Sub(t1))
}

func (e *Entangle) sections(w http.ResponseWriter, r *http.Request) {
	t1 := time.Now()
	cs, err := querySections(e.DB)
	if err != nil {
		http.Error(w, fmt.Sprintf("could not get sections: %s", err), http.StatusInternalServerError)
		return
	}
	if err := renderJSON(w, cs); err != nil {
		http.Error(w, "failed to encode", http.StatusInternalServerError)
		return
	}
	log.Printf("%d /sections returned successfully in %v", len(cs), time.Now().Sub(t1))
}

func renderJSON(w http.ResponseWriter, v interface{}) error {
	return json.NewEncoder(w).Encode(v)
}

func main() {
	var err error
	e := Entangle{Config: "config.yaml"}
	if err = e.startUp("local"); err != nil {
		log.Fatalf("failed to start api server: %s", err)
	}
	e.run(":8181")
}
