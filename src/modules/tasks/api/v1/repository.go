package v1

import (
	"check-list-be/src/modules/tasks/api/v1/models"
	"check-list-be/src/utils"
)

var repository = utils.NewRepository[models.Task]("tasks")

func Repository() utils.Repository[models.Task] {
	return repository
}
