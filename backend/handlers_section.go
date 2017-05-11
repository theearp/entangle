package main

import (
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2/bson"
)

func (e *Entangle) syncSections(w http.ResponseWriter, r *http.Request) {
	if err := syncSections(e.DB); err != nil {
		http.Error(w, fmt.Sprintf("failed to sync sections: %s", err), http.StatusInternalServerError)
		return
	}
	renderJSON(w, "success")
}

func (e *Entangle) sections(w http.ResponseWriter, r *http.Request) {
	var s []Section
	if err := e.DB.C("sections").Find(bson.M{}).All(&s); err != nil {
		http.Error(w, fmt.Sprintf("not found"), http.StatusInternalServerError)
		return
	}
	renderJSON(w, s)
}
