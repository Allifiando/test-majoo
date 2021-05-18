package productDomain

type Merchant struct {
	ID          int    `json:"id"`
	User        string `json:"user"`
	CompanyName string `json:"company_name"`
	Trademark   string `json:"trademark"`
	IsActive    bool   `json:"is_active"`
}

type SetMerchant struct {
	UserId      int    `json:"user_id" validate:"required" comment:"user id"`
	CompanyName string `json:"company_name" validate:"required" comment:"Nama perusahaan"`
	Trademark   string `json:"trademark" validate:"required" comment:"Trademark"`
	IsActive    bool   `json:"is_active,omitempty"`
}

type Outlet struct {
	ID         int    `json:"id"`
	MerchantId int    `json:"merchant_id"`
	Name       string `json:"name"`
	Location   string `json:"location"`
	IsActive   bool   `json:"is_active"`
}

type SetOutlet struct {
	MerchantId int    `json:"merchant_id"`
	Name       string `json:"name"`
	Location   string `json:"location"`
	IsActive   bool   `json:"is_active,omitempty"`
}

type Product struct {
	ID       int     `json:"id"`
	OutletId int     `json:"outlet_id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Qty      string  `json:"qty"`
	Filename string  `json:"filename"`
	IsActive bool    `json:"is_active"`
}

type SetProduct struct {
	OutletId int64   `form:"outlet_id" json:"outlet_id"`
	Name     string  `form:"name" json:"name"`
	Price    float64 `form:"price" json:"price"`
	Qty      string  `form:"qty" json:"qty"`
	Filename string  `form:"filename" json:"filename"`
	IsActive bool    `form:"is_active,omitempty"`
}
