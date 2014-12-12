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
  //*task.Description = "yaya"
  attribute := task.Attributes()[name]
  //panic(reflect.TypeOf(attribute))
  switch reflect.ValueOf(attribute) {
  case *int:
    *attribute.(*int) = value
  case *string:
    *attribute.(*string) = value
  }
//   "yaya"//reflect.ValueOf(value)
//  panic(attributes[name])
}

func (task *Task) Update(t *Task) {
  for name, attr := range t.Attributes() {
    if !reflect.ValueOf(attr).IsNil() {
      //panic(reflect.ValueOf(attr).Elem())
      task.Set(name, attr)
    }
  }
}

func (task *Task) Validate(errors *binding.Errors, req *http.Request) {
  if task.GetDescription() == "" {
    errors.Fields["description"] = "description is a required field"
  }
}

