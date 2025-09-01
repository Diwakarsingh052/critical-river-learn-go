package stores

import "app/internal/stores/models"

type UserRepository interface {
	Create(u models.User) (models.User, error)
	Delete(id int) bool
}
