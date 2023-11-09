package v1

import (
	m "check-list-be/src/middleware"
	"check-list-be/src/modules/users/api/v1/dto"
	v1 "check-list-be/src/modules/users/api/v1/models"

	"github.com/gofiber/fiber/v2"
)

func New() *fiber.App {
	v1.SyncIndexes()
	v1 := fiber.New()
	v1.Post("/", m.Validate[dto.CreateUserReq](m.Body), Create)

	return v1
}
