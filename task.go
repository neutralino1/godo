package main

import  (
  "reflect"
  "net/http"
  "github.com/codegangsta/martini-contrib/binding"
)

type Task struct {
  BaseModel `bson:",inline"`
  Description *string `bson:"description"`
  Position *int `bson:"position"`
  Done *bool `bson:"done"`
}

func (task *Task) Collection() string {
  return "tasks"
}

func (task *Task) Attributes() map[string]interface{} {
  return map[string]interface{}{
    "description": task.Description,
    "position": task.Position,
    "done": task.Done,
  }
}

func (task *Task) GetDescription() string {
  return *(task.Description)
}

func (task *Task) Set(name string, value interface{}) {
  attribute := task.Attributes()[name]
  switch attribute.(type) {
  default:
    *attribute.(*string) = *value.(*string)
  case *int:
    *attribute.(*int) = *value.(*int)
  case *string:
    *attribute.(*string) = *value.(*string)
  }
}

func (task *Task) Update(t *Task) {
  for name, attr := range t.Attributes() {
    if !reflect.ValueOf(attr).IsNil() {
      task.Set(name, attr)
    }
  }
}

func (task *Task) Validate(errors *binding.Errors, req *http.Request) {
  if task.GetDescription() == "" {
    errors.Fields["description"] = "description is a required field"
  }
}

