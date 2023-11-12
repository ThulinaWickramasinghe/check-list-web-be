package models

import (
	"check-list-be/src/database"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TaskStatus string

const (
	Todo TaskStatus = "todo"
	Done TaskStatus = "done"
)

type Task struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	UserID      primitive.ObjectID `bson:"user_id"`
	Description string             `json:"description" bson:"description,omitempty"`
	Status      TaskStatus         `json:"status" bson:"status,omitempty"`
	CreatedAt   string             `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt   string             `json:"updated_at" bson:"updated_at,omitempty"`
}

func (task Task) WithDefaults() Task {
	if task.Status == "" {
		task.Status = Todo
	}

	task.CreatedAt = time.Now().Format(time.RFC3339)
	task.UpdatedAt = time.Now().Format(time.RFC3339)

	return task
}

func SyncIndexes() {
	database.UseDefault().Collection("tasks").Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    bson.D{{Key: "_id", Value: -1}},
		Options: options.Index().SetUnique(true),
	})
}
