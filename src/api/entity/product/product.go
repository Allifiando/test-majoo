package productEntity

import (
	"context"
	"errors"
	productDomain "test-majoo/src/domain/product"
)

func (a *productEntity) GetListMerchantByUserId(c context.Context, offset, limit int, search, sort string, userId int64) (res []productDomain.Merchant, count int, err error) {
	ctx, cancel := context.WithTimeout(c, a.timeout)
	defer cancel()

	res, count, err = a.repo.GetListMerchantByUserId(ctx, offset, limit, search, sort, userId)
	if err != nil {
		err = errors.New("gagal get data merchant")
		return res, count, err
	}
	return
}

func (a *productEntity) GetListOutletByMerchantId(c context.Context, offset, limit int, search, sort string, merchantId int64) (res []productDomain.Outlet, count int, err error) {
	ctx, cancel := context.WithTimeout(c, a.timeout)
	defer cancel()

	res, count, err = a.repo.GetListOutletByMerchantId(ctx, offset, limit, search, sort, merchantId)
	if err != nil {
		err = errors.New("gagal get data outlet")
		return res, count, err
	}
	return
}

func (a *productEntity) GetListProductByOutletId(c context.Context, offset, limit int, search, sort string, outletId int64) (res []productDomain.Product, count int, err error) {
	ctx, cancel := context.WithTimeout(c, a.timeout)
	defer cancel()

	res, count, err = a.repo.GetListProductByOutletId(ctx, offset, limit, search, sort, outletId)
	if err != nil {
		err = errors.New("gagal get data merchant")
		return res, count, err
	}
	return
}

func (a *productEntity) GetMerchantById(c context.Context, id int64) (res productDomain.Merchant, err error) {
	ctx, cancel := context.WithTimeout(c, a.timeout)
	defer cancel()

	res, err = a.repo.GetMerchantById(ctx, id)
	if err != nil {
		err = errors.New("gagal get data merchant")
		return res, err
	}
	return
}

func (a *productEntity) GetMerchantByName(c context.Context, name string) (res productDomain.Merchant, err error) {
	ctx, cancel := context.WithTimeout(c, a.timeout)
	defer cancel()

	res, err = a.repo.GetMerchantByName(ctx, name)
	if err != nil {
		err = errors.New("gagal get data merchant")
		return res, err
	}
	return
}

func (a *productEntity) GetOutletById(c context.Context, id int64) (res productDomain.Outlet, err error) {
	ctx, cancel := context.WithTimeout(c, a.timeout)
	defer cancel()

	res, err = a.repo.GetOutletById(ctx, id)
	if err != nil {
		err = errors.New("gagal get data outlet")
		return res, err
	}
	return
}

func (a *productEntity) GetProductById(c context.Context, id int64) (res productDomain.Product, err error) {
	ctx, cancel := context.WithTimeout(c, a.timeout)
	defer cancel()

	res, err = a.repo.GetProductById(ctx, id)
	if err != nil {
		err = errors.New("gagal get data product")
		return res, err
	}
	return
}

func (a *productEntity) CreateMerchant(c context.Context, b productDomain.SetMerchant) (err error) {
	ctx, cancel := context.WithTimeout(c, a.timeout)
	defer cancel()

	err = a.repo.CreateMerchant(ctx, b)
	if err != nil {
		err = errors.New("gagal create merchant")
		return err
	}
	return
}

func (a *productEntity) CreateOutlet(c context.Context, b productDomain.SetOutlet) (err error) {
	ctx, cancel := context.WithTimeout(c, a.timeout)
	defer cancel()

	err = a.repo.CreateOutlet(ctx, b)
	if err != nil {
		err = errors.New("gagal create outlet")
		return err
	}
	return
}

func (a *productEntity) CreateProduct(c context.Context, b productDomain.SetProduct) (err error) {
	ctx, cancel := context.WithTimeout(c, a.timeout)
	defer cancel()

	err = a.repo.CreateProduct(ctx, b)
	if err != nil {
		err = errors.New("gagal create product")
		return err
	}
	return
}

func (a *productEntity) UpdateProduct(c context.Context, b productDomain.SetProduct, id int64) (err error) {
	ctx, cancel := context.WithTimeout(c, a.timeout)
	defer cancel()

	err = a.repo.UpdateProduct(ctx, b, id)
	if err != nil {
		err = errors.New("gagal update product")
		return err
	}
	return
}

func (a *productEntity) DeleteProduct(c context.Context, id int64) (err error) {
	ctx, cancel := context.WithTimeout(c, a.timeout)
	defer cancel()

	err = a.repo.DeleteProduct(ctx, id)
	if err != nil {
		err = errors.New("gagal delete product")
		return err
	}
	return
}
