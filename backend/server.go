package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

var (
	secrets *config
)

// Entangle represents the Entangle API service.
type Entangle struct {
	Config string
	DB     *mgo.Database
	Router *mux.Router
}

func (e *Entangle) startUp(env string) error {
	e.registerRoutes()
	var err error
	secrets, err = getConfig(e.Config)
	if err != nil {
		return fmt.Errorf("failed to get secrets: %s", err)
	}

	log.Println("initializing database...")
	s, err := mgo.Dial(secrets.MongoLocal.Address)
	if err != nil {
		return err
	}
	e.DB = s.DB(secrets.MongoLocal.DBName)
	return nil
}

func (e *Entangle) registerRoutes() {
	log.Println("initializing routes...")
	e.Router = mux.NewRouter()
	e.Router.HandleFunc("/", e.home).Methods("GET")

	// Sections
	e.Router.HandleFunc("/sections/sync", e.syncSections).Methods("GET")
	e.Router.HandleFunc("/sections", e.sections).Methods("GET")

	// Listings
	e.Router.HandleFunc("/listings", e.listings).Methods("GET", "UPDATE")
	e.Router.HandleFunc("/listing/{id}", e.getListing).Methods("GET")
	e.Router.HandleFunc("/listing/{id}/offerings", e.offerings).Methods("GET")
	e.Router.HandleFunc("/listing/{id}/images", e.images).Methods("GET")
	e.Router.HandleFunc("/listing/{id}/sync", e.syncListing).Methods("GET")
}

func (e *Entangle) run(addr string) {
	log.Println("Serving API")
	log.Fatal(http.ListenAndServe(addr, handlers.CORS()(e.Router)))
}

func (e *Entangle) home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Nothing here")
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
