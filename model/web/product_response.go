package web

import "time"

type ProductResponse struct {
	Id                int       `json:"id" example:"id"`
	ProductCategoryId int       `json:"product_category_id" example:"product category id"`
	SkuId             int       `json:"sku_id" example:"sku id"`
	Name              string    `json:"name" example:"product name"`
	Qty               int       `json:"qty" example:"qty"`
	QtyAvailable      int       `json:"qty_available" example:"qty available"`
	Type              int       `json:"type" example:"true = serial number, false = no serial number"`
	Price             int       `json:"price" example:"harga awal"`
	Sale              int       `json:"sale" example:"harga jual"`
	IsPpn             bool      `json:"is_ppn" example:"true = ppn, false = non ppn"`
	Active            bool      `json:"active" example:"true = active, false = inactive"`
	DeletedById       int       `json:"deleted_by_id" example:"10"`
	CreatedAt         time.Time `json:"created_at" example:"2025-04-03T01:13:12Z"`
	UpdatedAt         time.Time `json:"updated_at" example:"2025-04-03T01:13:12Z"`
	DeletedAt         time.Time `json:"deleted_at" example:"2025-04-03T01:13:12Z"`
}
