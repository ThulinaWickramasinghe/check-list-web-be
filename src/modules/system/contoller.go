package system

import (
	"check-list-be/src/global"
	"check-list-be/src/modules/system/dto"

	"github.com/gofiber/fiber/v2"
)

func Health(c *fiber.Ctx) error {
	return c.JSON(global.Response[*interface{}]{
		Message: "CheckList up and running",
	})
}

func Memory(c *fiber.Ctx) error {
	return c.JSON(global.Response[dto.MemStats]{
		Message: "Memory usage retrieved",
		Data:    GetMemoryUsage(),
	})
}
