package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

func (e *Entangle) offerings(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if !bson.IsObjectIdHex(id) {
		http.Error(w, fmt.Sprintf("invalid id: %s", id), http.StatusBadRequest)
		return
	}
	ref := bson.ObjectIdHex(id)
	var o []product
	if err := e.DB.C("products").Find(bson.M{"reference": ref}).All(&o); err != nil {
		http.Error(w, fmt.Sprintf("failed to fetch offerings: %s", err), http.StatusInternalServerError)
	}

	renderJSON(w, o)
}
