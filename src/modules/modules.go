package modules

import (
	"check-list-be/src/modules/auth"
	"check-list-be/src/modules/system"
	"check-list-be/src/modules/users"

	"github.com/gofiber/fiber/v2"
)

func New() *fiber.App {
	modules := fiber.New()

	modules.Mount("/users", users.New())
	modules.Mount("/auth", auth.New())
	modules.Mount("/system", system.New())

	return modules
}
