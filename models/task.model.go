package models

import (
	"github.com/sayedulkrm/go-curd-todo-2/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TodoList struct {
	Id     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Task   string             `json:"task"`
	Status string             `json:"status"`
}

var TodoCollection *mongo.Collection = database.OpenCollection(database.DBInstance(), "TodoList")
