package web

type ProductUpdateRequest struct {
	Id                int    `json:"id" example:"1"`
	ProductCategoryId int    `json:"product_category_id" example:"2"`
	SkuId             int    `json:"sku_id" example:"3"`
	Name              string `json:"name" example:"product test"`
	Qty               int    `json:"qty" example:"10"`
	QtyAvailable      int    `json:"qty_available" example:"5"`
	Type              int    `json:"type" example:"1"`
	Price             int    `json:"price" example:"1000"`
	Sale              int    `json:"sale"  example:"2000"`
	IsPpn             bool   `json:"is_ppn" example:"true"`
	Active            bool   `json:"active" example:"false"`
	DeletedById       int    `json:"deleted_by_id" example:"10"`
	// DeletedAt         time.Time `json:"deleted_at" example:"2023-10-01T00:00:00Z"`
}
