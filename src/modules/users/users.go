package users

import (
	"check-list-be/src/middleware"
	v1 "check-list-be/src/modules/users/api/v1"

	"github.com/gofiber/fiber/v2"
)

func New() *fiber.App {
	users := fiber.New()
	users.All("/*", middleware.AdminProtect)
	users.Mount("/v1", v1.New())

	return users
}
