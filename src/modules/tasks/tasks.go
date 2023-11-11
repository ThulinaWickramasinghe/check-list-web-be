package tasks

import (
	"check-list-be/src/middleware"
	v1 "check-list-be/src/modules/tasks/api/v1"

	"github.com/gofiber/fiber/v2"
)

func New() *fiber.App {
	tasks := fiber.New()
	tasks.All("/*", middleware.Protect)
	tasks.Mount("/v1", v1.New())

	return tasks
}
