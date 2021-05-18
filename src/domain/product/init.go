package productDomain

import "context"

type Entity interface {
	GetListMerchantByUserId(ctx context.Context, offset, limit int, search, sort string, userId int64) (res []Merchant, count int, err error)
	GetMerchantById(ctx context.Context, id int64) (res Merchant, err error)
	GetMerchantByName(ctx context.Context, name string) (res Merchant, err error)
	GetListOutletByMerchantId(ctx context.Context, offset, limit int, search, sort string, merchantId int64) (res []Outlet, count int, err error)
	GetOutletById(ctx context.Context, id int64) (res Outlet, err error)
	GetListProductByOutletId(ctx context.Context, offset, limit int, search, sort string, outletId int64) (res []Product, count int, err error)
	GetProductById(ctx context.Context, id int64) (res Product, err error)
	CreateMerchant(ctx context.Context, b SetMerchant) (err error)
	CreateOutlet(ctx context.Context, b SetOutlet) (err error)
	CreateProduct(ctx context.Context, b SetProduct) (err error)
	UpdateProduct(ctx context.Context, b SetProduct, id int64) (err error)
	DeleteProduct(ctx context.Context, id int64) (err error)
}
type Repo interface {
	GetListMerchantByUserId(ctx context.Context, offset, limit int, search, sort string, userId int64) (res []Merchant, count int, err error)
	GetMerchantById(ctx context.Context, id int64) (res Merchant, err error)
	GetMerchantByName(ctx context.Context, name string) (res Merchant, err error)
	GetListOutletByMerchantId(ctx context.Context, offset, limit int, search, sort string, merchantId int64) (res []Outlet, count int, err error)
	GetOutletById(ctx context.Context, id int64) (res Outlet, err error)
	GetListProductByOutletId(ctx context.Context, offset, limit int, search, sort string, outletId int64) (res []Product, count int, err error)
	GetProductById(ctx context.Context, id int64) (res Product, err error)
	CreateMerchant(ctx context.Context, b SetMerchant) (err error)
	CreateOutlet(ctx context.Context, b SetOutlet) (err error)
	CreateProduct(ctx context.Context, b SetProduct) (err error)
	UpdateProduct(ctx context.Context, b SetProduct, id int64) (err error)
	DeleteProduct(ctx context.Context, id int64) (err error)
}
