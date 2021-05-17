package userEntity

import (
	"time"

	userDomain "test-majoo/src/domain/user"
)

type UserEntity struct {
	repo    userDomain.Repo
	timeout time.Duration
}

func InitUserEntity(a userDomain.Repo, t time.Duration) userDomain.Entity {
	return &UserEntity{
		repo:    a,
		timeout: t,
	}
}
