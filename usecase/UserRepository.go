package usecase

import (
	"github.com/tsurusekazuki/clean-gin-gorm/app/domain"
)

type UserRepository interface {
	FindByID(id int) (event domain.Users, err error)
}