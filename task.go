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

func (task *Task) Collection() string {
  return "tasks"
}

func (task *Task) Fields() map[string]interface{} {
  return map[string]interface{}{
    "description": &task.Description,
    "position": &task.Position,
    "done": &task.Done,
  }
}

func (task *Task) GetId() bson.ObjectId {
  return task.Id
}

func (task *Task) Validate(errors *binding.Errors, req *http.Request) {
  if task.Description == "" {
    errors.Fields["description"] = "description is a required field"
  }
}

