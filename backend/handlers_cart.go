package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

func (e *Entangle) cart(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if bson.IsObjectIdHex(id) {
		var c shoppingCart
		if err := e.DB.C("carts").FindId(bson.ObjectIdHex(id)); err != nil {
			http.Error(w, fmt.Sprintf("failed to get cart: %s", err), http.StatusInternalServerError)
		}
		renderJSON(w, c)
	}
}
