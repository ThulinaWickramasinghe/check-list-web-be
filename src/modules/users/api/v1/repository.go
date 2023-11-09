package v1

import (
	"check-list-be/src/modules/users/api/v1/models"
	"check-list-be/src/utils"
)

var repository = utils.NewRepository[models.User]("users")

func Repository() utils.Repository[models.User] {
	return repository
}
