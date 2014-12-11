package main

import  (
  "net/http"
  "labix.org/v2/mgo/bson"
  "github.com/codegangsta/martini-contrib/binding"
)

type Task struct {
  Id bson.ObjectId `bson:"_id"`
  Description string `bson:"description"`
  Position int `bson:"position"`
  Done bool `bson:"done"`
}

func (taskAttr *Task) Validate(errors *binding.Errors, req *http.Request) {
  if taskAttr.Description == "" {
    errors.Fields["description"] = "description is a required field"
  }
}