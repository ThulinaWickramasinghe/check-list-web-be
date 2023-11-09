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

type UserRole string

// TODO: Remove Roles once the project is complete
const (
	Admin   UserRole = "admin"
	Regular UserRole = "regular"
)

type User struct {
	ID               primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name             string             `json:"name" bson:"name,omitempty"`
	Email            string             `json:"email" bson:"email,omitempty"`
	Password         string             `json:"password" bson:"password,omitempty"`
	Verified         bool               `json:"verified" bson:"verified"`
	VerificationCode *string            `json:"verification_code" bson:"verification_code,omitempty"`
	Role             UserRole           `json:"role" bson:"role,omitempty"` //TODO: Remove once project is complete
	CreatedAt        string             `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt        string             `json:"updated_at" bson:"updated_at,omitempty"`
}

func (u User) WithDefaults() User {
	// TODO: Remove once the project is complete
	if u.Role == "" {
		u.Role = Regular
	}

	u.CreatedAt = time.Now().Format(time.RFC3339)
	u.UpdatedAt = time.Now().Format(time.RFC3339)

	return u
}

func (u User) Secure() User {
	u.Password = ""
	return u
}

func SyncIndexes() {
	database.UseDefault().Collection("users").Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: -1}},
		Options: options.Index().SetUnique(true),
	})
}
