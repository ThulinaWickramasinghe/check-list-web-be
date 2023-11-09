package v1

import (
	"check-list-be/src/modules/users/api/v1/dto"
	"check-list-be/src/modules/users/api/v1/models"
	"check-list-be/src/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"github.com/sethvargo/go-password/password"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func createUser(c *fiber.Ctx, payload dto.CreateUserReq) *dto.CreateUserRes {
	log.Info("Creating a user within system - ", payload.Email)

	verificationCode := uuid.New().String()
	generatedPassword, _ := password.Generate(8, 2, 1, false, false)

	insertedId := repository.Create(models.User{
		Email:            payload.Email,
		Name:             payload.Name,
		VerificationCode: &verificationCode,
		Password:         utils.HashStr(generatedPassword),
		Role:             models.UserRole(payload.Role),
	}.WithDefaults())

	return &dto.CreateUserRes{
		ID:       insertedId,
		Password: generatedPassword,
	}
}

func getUser(c *fiber.Ctx, id primitive.ObjectID) *dto.GetUserRes {
	log.Info("Fetching a user with ID - " + id.String())

	user := repository.FindByID(id)

	return &dto.GetUserRes{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      string(user.Role),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
