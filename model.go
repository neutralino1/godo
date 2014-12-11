package main

import "labix.org/v2/mgo/bson"

type Model interface {
  GetId() bson.ObjectId
  Collection() string
}