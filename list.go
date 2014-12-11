package main

import  (
  "net/http"
  "labix.org/v2/mgo/bson"
  "github.com/codegangsta/martini-contrib/binding"
)

type List struct {
  Id bson.ObjectId `bson:"_id"`
  Name string `bson:"name"`
  Tasks []Task `bson:"tasks"`
}

func (list *List) Collection() string {
  return "lists"
}

func (list *List) GetId() bson.ObjectId {
  return list.Id
}

func (listAttr *List) Validate(errors *binding.Errors, req *http.Request) {
  if listAttr.Name == "" {
    errors.Fields["name"] = "name is a required field"
  }
}