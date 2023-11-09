package v1

import (
	"check-list-be/src/global"
	"check-list-be/src/modules/users/api/v1/dto"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Create(c *fiber.Ctx) error {
	payload := new(dto.CreateUserReq)
	c.BodyParser(payload)
	res := createUser(c, *payload)

	return c.Status(fiber.StatusCreated).JSON(global.Response[dto.CreateUserRes]{
		Message: "User created successfully",
		Data:    res,
	})
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")

	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is a hex string")
	}

	res := getUser(c, objectID)

	return c.Status(200).JSON(global.Response[dto.GetUserRes]{
		Message: "User retrieved",
		Data:    res,
	})
}
