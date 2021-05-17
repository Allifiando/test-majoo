package userEntity

import (
	"context"
	"errors"
	userDomain "test-majoo/src/domain/user"
)

func (a *UserEntity) GetListUser(c context.Context, offset, limit int, search, sort string) (res []userDomain.User, count int, err error) {
	ctx, cancel := context.WithTimeout(c, a.timeout)
	defer cancel()

	res, count, err = a.repo.GetListUser(ctx, offset, limit, search, sort)
	if err != nil {
		err = errors.New("gagal get data user")
		return res, count, err
	}
	return
}

func (a *UserEntity) Login(c context.Context, b userDomain.Login) (res userDomain.UserLogin, err error) {
	ctx, cancel := context.WithTimeout(c, a.timeout)
	defer cancel()

	res, err = a.repo.Login(ctx, b)
	if err != nil {
		err = errors.New("gagal login")
		return res, err
	}
	return
}

func (a *UserEntity) GetUserById(c context.Context, id int) (res userDomain.User, err error) {
	ctx, cancel := context.WithTimeout(c, a.timeout)
	defer cancel()

	res, err = a.repo.GetUserById(ctx, id)
	if err != nil {
		err = errors.New("gagal get data user")
		return res, err
	}

	return
}

func (a *UserEntity) Create(c context.Context, b userDomain.SetUser) (err error) {
	ctx, cancel := context.WithTimeout(c, a.timeout)
	defer cancel()

	err = a.repo.Create(ctx, b)
	if err != nil {
		err = errors.New("gagal create user")
		return err
	}
	return
}

func (a *UserEntity) Update(c context.Context, b userDomain.SetUser, id int) (err error) {
	ctx, cancel := context.WithTimeout(c, a.timeout)
	defer cancel()

	err = a.repo.Update(ctx, b, id)
	if err != nil {
		err = errors.New("gagal update user")
		return err
	}
	return
}

func (a *UserEntity) Delete(c context.Context, id int) (err error) {
	ctx, cancel := context.WithTimeout(c, a.timeout)
	defer cancel()

	err = a.repo.Delete(ctx, id)
	if err != nil {
		err = errors.New("gagal delete user")
		return err
	}
	return
}
