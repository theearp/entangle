package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

func (e *Entangle) getListing(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if !bson.IsObjectIdHex(id) {
		http.Error(w, fmt.Sprintf("invalid id: %s", id), http.StatusBadRequest)
		return
	}
	oid := bson.ObjectIdHex(id)
	var l listing
	if err := e.DB.C("listings").FindId(oid).One(&l); err != nil {
		http.Error(w, fmt.Sprintf("failed to fetch listing: %s", err), http.StatusInternalServerError)
		return
	}
	renderJSON(w, l)
}

func (e *Entangle) listings(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		var l []listing
		if err := e.DB.C("listings").Find(bson.M{}).All(&l); err != nil {
			http.Error(w, fmt.Sprintf("failed to fetch listing: %s", err), http.StatusInternalServerError)
			return
		}
		renderJSON(w, l)
	case "UPDATE":
		if err := fetchAndWriteActiveListings(e.DB); err != nil {
			http.Error(w, fmt.Sprintf("failed to update listings: %s", err), http.StatusInternalServerError)
			return
		}
		renderJSON(w, "success")
	default:
		http.Error(w, "Not Implemented", http.StatusNotImplemented)
		return
	}
}

func (e *Entangle) syncListing(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if !bson.IsObjectIdHex(id) {
		http.Error(w, fmt.Sprintf("invalid id: %s", id), http.StatusBadRequest)
		return
	}
	oid := bson.ObjectIdHex(id)
	var l listing
	if err := e.DB.C("listings").FindId(oid).One(&l); err != nil {
		http.Error(w, fmt.Sprintf("failed to fetch listing: %s", err), http.StatusInternalServerError)
		return
	}
	inventory, err := l.getInventoryRemote()
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to fetch inventory: %s", err), http.StatusInternalServerError)
		return
	}
	for _, product := range inventory {
		product.ID = bson.NewObjectId()
		product.Reference = l.ID
		if err := e.DB.C("products").Insert(product); err != nil {
			log.Printf("failed to insert product %q: %s", product.ProductID, err)
		}
	}

	// Write image data.
	images, err := l.getImagesRemote()
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to fetch images: %s", err), http.StatusInternalServerError)
		return
	}
	for _, i := range images {
		i.ID = bson.NewObjectId()
		i.Reference = l.ID
		if err := e.DB.C("images").Insert(i); err != nil {
			log.Printf("failed to insert image data for %q: %s", i.ListingID, err)
		}
	}
	renderJSON(w, "success")
}
