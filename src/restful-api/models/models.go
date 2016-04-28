package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type User struct {
	Id           bson.ObjectId `bson:"_id,omitempty" json:"id"`
	FirstName    string        `json:"firstname"`
	LastName     string        `json:"lastname"`
	Email        string        `json:"email"`
	Password     string        `json:"password,omitempty"`
	HashPassword string        `json:"hashpassword,omitempty"`
}

type Task struct {
	Id          bson.ObjectId `bson:"_id,omitempty" json:"id"`
	CreatedBy   string        `json:"createdby"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	CreatedOn   time.Time     `json:"createon,omitempty"`
	Due         time.Time     `json:"due,omitempty"`
	Status      string        `json:"status,omitempty"`
	Tags        string        `json:"tags,omitempty"`
}

type TaskNote struct {
	Id          bson.ObjectId `bson:"_id,omitempty" json:"id"`
	TaskId      bson.ObjectId `json:"taskid"`
	Description string        `json:"description"`
	CreatedOn   time.Time     `json:"createdon,omitempty"`
}
