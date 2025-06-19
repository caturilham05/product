package service

import (
	"caturilham05/product/exception"
	"caturilham05/product/helper"
	"caturilham05/product/model/domain"
	"caturilham05/product/model/web"
	"caturilham05/product/repository"
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
)

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

// FindByName implements ProductService.
func (p *ProductServiceImpl) FindByName(ctx context.Context, productName string) web.ProductResponse {
	tx, err := p.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product, err := p.ProductRepository.FindByName(ctx, tx, productName)
	if err != nil {
		if err.Error() == "produk tidak ditemukan" {
			panic(exception.NewNotFoundError(err.Error()))
		} else {
			helper.PanicIfError(err)
		}
	}

	return helper.ToProductResponse(product)
}

// Create implements ProductService.
func (p *ProductServiceImpl) Create(ctx context.Context, request web.ProductCreateRequest) web.ProductResponse {
	err := p.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := p.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	productCheck, err := p.ProductRepository.FindByName(ctx, tx, request.Name)
	if err != nil {
		helper.PanicIfError(err)
	}

	if productCheck.Id != 0 {
		panic(exception.NewBadRequestError("produk dengan nama " + request.Name + " sudah terdaftar"))
	}

	product := domain.Product{
		ProductCategoryId: request.ProductCategoryId,
		SkuId:             request.SkuId,
		Name:              request.Name,
		Qty:               request.Qty,
		QtyAvailable:      request.QtyAvailable,
		Type:              request.Type,
		Price:             request.Price,
		Sale:              request.Sale,
		IsPpn:             request.IsPpn,
		Active:            request.Active,
	}

	product = p.ProductRepository.Save(ctx, tx, product)
	if product.Id == 0 {
		panic(exception.NewBadRequestError("gagal membuat produk"))
	}
	return helper.ToProductResponse(product)
}

// Delete implements ProductService.
func (p *ProductServiceImpl) Delete(ctx context.Context, productId int, userId int) {
	tx, err := p.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product, err := p.ProductRepository.FindById(ctx, tx, productId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	p.ProductRepository.Delete(ctx, tx, product.Id, userId)
}

// FindAll implements ProductService.
func (p *ProductServiceImpl) FindAll(ctx context.Context) []web.ProductResponse {
	tx, err := p.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	products := p.ProductRepository.FindAll(ctx, tx)
	return helper.ToProductResponses(products)
}

// FindById implements ProductService.
func (p *ProductServiceImpl) FindById(ctx context.Context, productId int) web.ProductResponse {
	tx, err := p.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product, err := p.ProductRepository.FindById(ctx, tx, productId)
	if err != nil {
		if err.Error() == "produk tidak ditemukan" {
			panic(exception.NewNotFoundError(err.Error()))
		} else {
			helper.PanicIfError(err)
		}
	}

	return helper.ToProductResponse(product)
}

// Update implements ProductService.
func (p *ProductServiceImpl) Update(ctx context.Context, request web.ProductUpdateRequest) web.ProductResponse {
	err := p.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := p.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// check product exist
	product, err := p.ProductRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	product.ProductCategoryId = request.ProductCategoryId
	product.SkuId = request.SkuId
	product.Name = request.Name
	product.Qty = request.Qty
	product.QtyAvailable = request.QtyAvailable
	product.Type = request.Type
	product.Price = request.Price
	product.Sale = request.Sale
	product.IsPpn = request.IsPpn
	product.Active = request.Active
	product.DeletedById = request.DeletedById

	product = p.ProductRepository.Update(ctx, tx, product)
	return helper.ToProductResponse(product)
}

func NewProductService(productRepository repository.ProductRepository, DB *sql.DB, validate *validator.Validate) ProductService {
	return &ProductServiceImpl{
		ProductRepository: productRepository,
		DB:                DB,
		Validate:          validate,
	}
}
