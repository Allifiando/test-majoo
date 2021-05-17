package productDomain

type Merchant struct {
	ID     int    `json:"id"`
	UserId int    `json:"user_id"`
	Name   string `json:"name"`
}

type SetMerchant struct {
	UserId int    `json:"user_id"`
	Name   string `json:"name"`
}

type Outlet struct {
	ID         int    `json:"id"`
	MerchantId int    `json:"merchant_id"`
	Name       string `json:"name"`
	Location   string `json:"location"`
}

type SetOutlet struct {
	MerchantId int    `json:"merchant_id"`
	Name       string `json:"name"`
	Location   string `json:"location"`
}

type Product struct {
	ID       int     `json:"id"`
	OutletId int     `json:"outlet_id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Qty      string  `json:"qty"`
	Filename string  `json:"filename"`
}

type SetProduct struct {
	OutletId int     `json:"outlet_id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Qty      string  `json:"qty"`
	Filename string  `json:"filename"`
}
