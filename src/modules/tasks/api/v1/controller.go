package v1

import (
	"check-list-be/src/global"
	"check-list-be/src/modules/tasks/api/v1/dto"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Create(c *fiber.Ctx) error {
	payload := new(dto.CreateTaskReq)
	c.BodyParser(payload)
	res := createTask(c, *payload)

	return c.Status(fiber.StatusCreated).JSON(global.Response[dto.CreateTaskRes]{
		Message: "Task created successfully",
		Data:    res,
	})
}

func GetTask(c *fiber.Ctx) error {
	id := c.Params("id")

	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is a hex string")
	}

	res := getTask(c, objectID)

	return c.Status(200).JSON(global.Response[dto.GetTaskRes]{
		Message: "Task retrieved",
		Data:    res,
	})
}

func GetTasks(c *fiber.Ctx) error {

	res := getTasks(c)

	return c.Status(200).JSON(global.Response[[]dto.GetTaskRes]{
		Message: "Task retrieved",
		Data:    res,
	})
}
