package userEntity

import (
	"time"

	userDomain "test-majoo/src/domain/user"
)

type userEntity struct {
	repo    userDomain.Repo
	timeout time.Duration
}

func InitUserEntity(a userDomain.Repo, t time.Duration) userDomain.Entity {
	return &userEntity{
		repo:    a,
		timeout: t,
	}
}
