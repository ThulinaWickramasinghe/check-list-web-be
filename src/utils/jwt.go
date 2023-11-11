package utils

import (
	"check-list-be/src/config"
	"check-list-be/src/modules/users/api/v1/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GenerateUserJWTToken(user models.User, refresh bool) string {
	expiry := time.Hour * 24

	if refresh {
		expiry = time.Hour * 48
	}

	claims := jwt.MapClaims{
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(expiry).Unix(),
		"data": user,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(config.Env.JWTSecret))

	if err != nil {
		panic(fiber.NewError(fiber.StatusInternalServerError, "Error genreating jwt token"))
	}

	return t
}

func ValidateUserJWTTOken(token string) *models.User {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Env.JWTSecret), nil
	})

	if err != nil || !parsedToken.Valid {
		panic(fiber.NewError(fiber.StatusUnauthorized, "Invalid token"))
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		panic(fiber.NewError(fiber.StatusUnauthorized, "Invalid token"))
	}

	claimsData, ok := claims["data"].(map[string]interface{})
	if !ok {
		panic(fiber.NewError(fiber.StatusUnauthorized, "Invalid data format in claims"))
	}

	objID, err := primitive.ObjectIDFromHex(claimsData["_id"].(string))
	role, ok := claimsData["role"].(string)
	verificationCode, ok := claimsData["verification_code"].(string)

	user := models.User{
		Email:            claimsData["email"].(string),
		Role:             models.UserRole(role),
		Name:             claimsData["name"].(string),
		ID:               objID,
		Password:         claimsData["password"].(string),
		CreatedAt:        claimsData["created_at"].(string),
		VerificationCode: &verificationCode,
		Verified:         claimsData["verified"].(bool),
	}

	return &user
}
