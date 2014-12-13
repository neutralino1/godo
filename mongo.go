package main

import (
  "github.com/codegangsta/martini"
  "labix.org/v2/mgo"
  "labix.org/v2/mgo/bson"
)

func mongo() martini.Handler {
 session, err := mgo.Dial("localhost/godo")
  if err != nil {
    panic(err)
  }

  return func (c martini.Context) {
    reqSession := session.Clone()
    c.Map(&Database{Conn: reqSession.DB("godo")})
    defer reqSession.Close()
    c.Next()
  }
}

type Model interface {
  GetId() bson.ObjectId
  Collection() string
  Attributes() map[string]interface{}
}

type Database struct {
  Conn *mgo.Database
}

func (db *Database) Find(m Model, id string) error {
  return (*db).Conn.C(m.Collection()).FindId(bson.ObjectIdHex(id)).One(m)
}

func (db *Database) All(m Model, models interface{}) error {  
  return (*db).Conn.C(m.Collection()).Find(bson.M{}).All(models)
}

func (db *Database) Insert(m Model) error {
  return (*db).Conn.C(m.Collection()).Insert(m)
}

func (db *Database) Update(m Model) error {
  return (*db).Conn.C(m.Collection()).Update(bson.M{"_id": m.GetId()}, m)
}

func (db *Database) UpdateSub(doc Model, docId string, sub Model) error {
  set := bson.M{}
  for k, v := range sub.Attributes() {
    set[sub.Collection() + ".$." + k] = v
  }
  return (*db).Conn.C(doc.Collection()).Update(bson.M{"_id": bson.ObjectIdHex(docId), sub.Collection() + "._id": sub.GetId()}, bson.M{"$set": set})
}