package helper

import (
	"caturilham05/product/model/domain"
	"caturilham05/product/model/web"
)

func ToProductResponse(product domain.Product) web.ProductResponse {
	return web.ProductResponse{
		Id:                product.Id,
		ProductCategoryId: product.ProductCategoryId,
		SkuId:             product.SkuId,
		Name:              product.Name,
		Qty:               product.Qty,
		QtyAvailable:      product.QtyAvailable,
		Type:              product.Type,
		Price:             product.Price,
		Sale:              product.Sale,
		IsPpn:             product.IsPpn,
		Active:            product.Active,
		DeletedById:       product.DeletedById,
		CreatedAt:         product.CreatedAt,
		UpdatedAt:         product.UpdatedAt,
		DeletedAt:         product.DeletedAt,
	}

}

func ToProductResponses(products []domain.Product) []web.ProductResponse {
	var productResponses []web.ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, ToProductResponse(product))
	}
	return productResponses
}
