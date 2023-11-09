package middleware

import (
	"check-list-be/src/global"
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/mongo"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	message := "Patching things up. This'll be over in no time."

	var e *fiber.Error

	if errors.As(err, &e) {
		code = e.Code
		message = e.Message
	}

	if mongo.IsDuplicateKeyError(err) {
		code = fiber.StatusBadRequest
		message = "Resource already exists"
	}

	log.Error("Request error: ", err.Error())

	return ctx.Status(code).JSON(global.Response[*interface{}]{
		Message: message,
	})
}
