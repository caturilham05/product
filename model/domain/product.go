package domain

import "time"

type Product struct {
	Id                int
	ProductCategoryId int
	SkuId             int
	Name              string
	Qty               int
	QtyAvailable      int
	Type              int
	Price             int
	Sale              int
	IsPpn             bool
	Active            bool
	DeletedById       int
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         time.Time
}
