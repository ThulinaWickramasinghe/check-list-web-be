package v1

import (
	"check-list-be/src/modules/tasks/api/v1/dto"
	t "check-list-be/src/modules/tasks/api/v1/models"
	u "check-list-be/src/modules/users/api/v1/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func createTask(c *fiber.Ctx, payload dto.CreateTaskReq) *dto.CreateTaskRes {
	log.Info("Creating a task within system - ")

	user := c.Locals("user").(*u.User)

	insertedId := repository.Create(t.Task{
		Description: payload.Description,
		Status:      t.TaskStatus(payload.Status),
		UserID:      user.ID,
	}.WithDefaults())

	if payload.Status == "" {
		payload.Status = "todo"
	}

	return &dto.CreateTaskRes{
		ID:          insertedId,
		Description: payload.Description,
		Status:      t.TaskStatus(payload.Status),
	}
}

func getTask(c *fiber.Ctx, id primitive.ObjectID) *dto.GetTaskRes {
	log.Info("Fetching a task with ID - " + id.String())

	user := c.Locals("user").(*u.User)

	task := repository.FindByID(id)

	if task.UserID != user.ID {
		panic(fiber.NewError(fiber.StatusNotFound, "No tasks match the given ID"))
	}

	return &dto.GetTaskRes{
		ID:          task.ID,
		Description: task.Description,
		Status:      task.Status,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}
}

func getTasks(c *fiber.Ctx) *[]dto.GetTaskRes {
	user := c.Locals("user").(*u.User)

	log.Info("Fetching all tasks of ", user.ID)

	filter := map[string]interface{}{"user_id": user.ID}
	tasks := repository.FindAllWithFilter(filter)

	var taskResponses []dto.GetTaskRes

	for _, task := range tasks {
		taskResponses = append(taskResponses, dto.GetTaskRes{
			ID:          task.ID,
			Description: task.Description,
			Status:      task.Status,
			CreatedAt:   task.CreatedAt,
			UpdatedAt:   task.UpdatedAt,
		})
	}

	return &taskResponses
}
