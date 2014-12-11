package main

import "labix.org/v2/mgo/bson"

type BaseModel struct {
  Id bson.ObjectId `bson:"_id"`
}

func (base *BaseModel) GetId() bson.ObjectId {
  return base.Id
}