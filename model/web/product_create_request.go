package web

type ProductCreateRequest struct {
	Name              string `json:"name" example:"product name"`
	ProductCategoryId int    `json:"product_category_id" example:"product category id"`
	SkuId             int    `json:"sku_id" example:"sku id"`
	Qty               int    `json:"qty" example:"qty"`
	QtyAvailable      int    `json:"qty_available" example:"qty available"`
	Type              int    `json:"type" example:"true = serial number, false = no serial number"`
	Price             int    `json:"price" example:"harga awal"`
	Sale              int    `json:"sale" example:"harga jual"`
	IsPpn             bool   `json:"is_ppn" example:"true = ppn, false = non ppn"`
	Active            bool   `json:"active" example:"true = active, false = inactive"`
}
