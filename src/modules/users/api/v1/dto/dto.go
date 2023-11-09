package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type CreateUserReq struct {
	Name  string `validate:"required"`
	Email string `validate:"required,email"`
	Role  string
}

type CreateUserRes struct {
	ID       primitive.ObjectID `json:"_id"`
	Password string             `json:"password"`
}

type GetUserRes struct {
	ID        primitive.ObjectID `json:"_id"`
	Name      string             `json:"name" validate:"required"`
	Email     string             `json:"email" validate:"required,email"`
	Role      string             `json:"role"`
	CreatedAt string             `json:"created_at"`
	UpdatedAt string             `json:"updated_at"`
}
