package main

import "gopkg.in/mgo.v2/bson"

type shoppingCart struct {
	ID    bson.ObjectId `bson:"_id" json:"id"`
	Items []listing
}
