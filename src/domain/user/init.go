package userDomain

import "context"

type Entity interface {
	GetListUser(ctx context.Context, offset, limit int, search, sort string) (res []User, count int, err error)
	GetUserById(ctx context.Context, id int) (res User, err error)
	Login(ctx context.Context, b Login) (res UserLogin, err error)
	Create(ctx context.Context, b SetUser) (err error)
	Update(ctx context.Context, b SetUser, id int) (err error)
	Delete(ctx context.Context, id int) (err error)
}
type Repo interface {
	GetListUser(ctx context.Context, offset, limit int, search, sort string) (res []User, count int, err error)
	GetUserById(ctx context.Context, id int) (res User, err error)
	Login(ctx context.Context, b Login) (res UserLogin, err error)
	Create(ctx context.Context, b SetUser) (err error)
	Update(ctx context.Context, b SetUser, id int) (err error)
	Delete(ctx context.Context, id int) (err error)
}
