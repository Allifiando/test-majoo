package productEntity

import (
	"time"

	productDomain "test-majoo/src/domain/product"
)

type productEntity struct {
	repo    productDomain.Repo
	timeout time.Duration
}

func InitProductEntity(a productDomain.Repo, t time.Duration) productDomain.Entity {
	return &productEntity{
		repo:    a,
		timeout: t,
	}
}
