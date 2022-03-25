package repository

import (
	"go-baseline/entity"
)

type UserRepository interface {
	FindAll() (users []entity.User)

	FindByUsername(username string) (user entity.User, err error)

	Create(user entity.User)

	Update(user entity.User)

	Delete(id string)
}
