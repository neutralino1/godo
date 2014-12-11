package main

import  (
  "net/http"
  "github.com/codegangsta/martini-contrib/binding"
)

type Task struct {
  BaseModel `bson:",inline"`
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

func (task *Task) Validate(errors *binding.Errors, req *http.Request) {
  if task.Description == "" {
    errors.Fields["description"] = "description is a required field"
  }
}

