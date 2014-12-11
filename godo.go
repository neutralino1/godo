package main

import (
  "net/http"
  "github.com/codegangsta/martini"
  "github.com/codegangsta/martini-contrib/binding"
  "encoding/json"
  "labix.org/v2/mgo"
  "labix.org/v2/mgo/bson"
)

func main() {
  m := martini.Classic()
  m.Use(setContentType())
  m.Use(mongo())

  m.Get("/", home)
  m.Get("/lists", getLists)
  m.Get("/lists/:id", getList)
  m.Post("/lists", binding.Json(List{}), createList)
  
  m.Get("/tasks/:id", getTask)
  m.Post("/lists/:listId/tasks", binding.Json(Task{}), createTask)
  m.Put("/lists/:listId/tasks/:id", binding.Json(Task{}), updateTask)

  m.Run()
}

type errorMsg struct {
  Msg string `json:"msg"`
}

type jsonConvertible interface {}

func jsonString(obj jsonConvertible) (s string) {
  jsonObj, err := json.Marshal(obj)

  if err != nil {
    s = ""
  } else {
    s = string(jsonObj)
  }
  return
}

func setContentType() martini.Handler {
  return func (writer http.ResponseWriter) {
    writer.Header().Set("Content-Type", "application/json")
  }
}

func home() (int, string) {
  return http.StatusOK, "Now go do gogo!!"
}

func getTask(params martini.Params, db *mgo.Database) (int, string) {
  id := params["id"]
  return http.StatusOK, id
}

func createTask(taskAttr Task, err binding.Errors, params martini.Params, db *Database) (int, string) {
  listId := params["listId"]
  list := List{}
  if db.Find(&list, listId) != nil {
    return http.StatusNotFound, jsonString(errorMsg{"No list found with id " + listId})
  }
  if err.Count() > 0 {
    return http.StatusConflict, jsonString(errorMsg{err.Overall["description"]})
  }
  taskAttr.Id = bson.NewObjectId()
  list.Tasks = append(list.Tasks, taskAttr)
  if dbErr := db.Update(&list) ; dbErr != nil {
    panic(dbErr)
  }
  return http.StatusOK, jsonString(list)
}

func updateTask(taskAttr Task, err binding.Errors, params martini.Params, db *mgo.Database) (int, string) {
  if err.Count() > 0 {
    return http.StatusConflict, jsonString(errorMsg{err.Overall["missing-requirement"]})
  }
  listId := params["listId"]
  taskAttr.Id = bson.ObjectIdHex(params["id"])
  if dbErr := db.C("lists").Update(bson.M{"_id": bson.ObjectIdHex(listId), "tasks._id": taskAttr.Id}, bson.M{"$set": bson.M{"tasks.$.description": taskAttr.Description}}) ; dbErr != nil {
    return http.StatusConflict, jsonString(dbErr)//errorMsg{"No list with id " + listId})
  }

  return http.StatusOK, jsonString(taskAttr)
}

func getLists(db *Database) (int, string) {
  lists := []List{}
  if err := db.All(new(List), &lists) ; err != nil {
    panic(err)
  }
  return http.StatusOK, jsonString(lists)
}

func getList(params martini.Params, db *Database) (int, string) {
  id := params["id"]
  list := List{}
  if db.Find(&list, id) != nil {
    return http.StatusNotFound, jsonString(errorMsg{"No list found with id " + id})
  }
  return http.StatusOK, jsonString(list)
}

func createList(listAttr List, err binding.Errors, db *Database) (int, string) {
  if err.Count() > 0 {
    return http.StatusConflict, jsonString(errorMsg{err.Overall["missing-requirement"]})
  }
  listAttr.Id = bson.NewObjectId()
  dbErr := db.Insert(&listAttr) ; if dbErr != nil {
    return http.StatusConflict, jsonString(dbErr)
  }
  return http.StatusOK, jsonString(listAttr)
}