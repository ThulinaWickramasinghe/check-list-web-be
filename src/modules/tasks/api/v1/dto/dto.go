package dto

import (
	"check-list-be/src/modules/tasks/api/v1/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateTaskReq struct {
	Description string            `json:"description" bson:"description,omitempty"`
	Status      models.TaskStatus `json:"status" bson:"status,omitempty"`
}

type CreateTaskRes struct {
	ID          primitive.ObjectID `json:"_id"`
	Description string             `json:"description" bson:"description,omitempty"`
	Status      models.TaskStatus  `json:"status" bson:"status,omitempty"`
}

type GetTaskRes struct {
	ID          primitive.ObjectID `json:"_id"`
	Description string             `json:"description" bson:"description,omitempty"`
	Status      models.TaskStatus  `json:"status" bson:"status,omitempty"`
	CreatedAt   string             `json:"created_at"`
	UpdatedAt   string             `json:"updated_at"`
}

type ToggleStatusRes struct {
	ID     primitive.ObjectID `json:"_id"`
	Status models.TaskStatus  `json:"status" bson:"status,omitempty"`
}

type UpdateTaskDescReq struct {
	Description string `json:"description" bson:"description,omitempty"`
}

type UpdateTaskDescRes struct {
	ID          primitive.ObjectID `json:"_id"`
	Description string             `json:"description" bson:"description,omitempty"`
}
