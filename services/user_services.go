package services

import (
	"github.com/manuelhdez/bs_users-api/domain/users"
	"github.com/manuelhdez/bs_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil
}
