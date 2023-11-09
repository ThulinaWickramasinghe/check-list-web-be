package middleware

import (
	"check-list-be/src/modules/users/api/v1/models"
	"check-list-be/src/utils"

	"github.com/gofiber/fiber/v2"
)

func Protect(ctx *fiber.Ctx) error {
	token := ctx.Get(fiber.HeaderAuthorization)

	if "" == token {
		panic(fiber.NewError(fiber.StatusUnauthorized, "Missing bearer token"))
	}

	user := utils.ValidateUserJWTTOken(token[len("Bearer "):])
	ctx.Locals("user", user)

	return ctx.Next()
}

// TODO: Remove once the project is complete
func AdminProtect(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*models.User)
	if user.Role != models.Admin {
		panic(fiber.NewError(fiber.StatusUnauthorized, "You are not authorized to access this resource"))
	}
	return ctx.Next()
}
