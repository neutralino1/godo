package main

import  (
  "net/http"
  "github.com/codegangsta/martini-contrib/binding"
)

type List struct {
  BaseModel `bson:",inline"`
  Name string `bson:"name"`
  Tasks []Task `bson:"tasks"`
}

func (list *List) Collection() string {
  return "lists"
}

func (list *List) Fields() map[string]interface{} {
  return map[string]interface{}{}
}

func (listAttr *List) Validate(errors *binding.Errors, req *http.Request) {
  if listAttr.Name == "" {
    errors.Fields["name"] = "name is a required field"
  }
}