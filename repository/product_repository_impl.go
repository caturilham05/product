package repository

import (
	"caturilham05/product/helper"
	"caturilham05/product/model/domain"
	"context"
	"database/sql"
	"time"

	"errors"
)

type ProductRepositoryImpl struct {
}

// Delete implements ProductRepository.
func (p *ProductRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, productId int, userId int) {
	SQL := "UPDATE `product` SET `deleted_at` = ?, `deleted_by_id` = ?, `active` = ? WHERE `id` = ?"
	// timeNow := time.Now().UTC()
	_, err := tx.ExecContext(ctx, SQL, time.Now(), userId, 0, productId)
	helper.PanicIfError(err)
}

// FIndAll implements ProductRepository.
func (p *ProductRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Product {
	SQL := "SELECT `id`, `product_category_id`, `sku_id`, `name`, `qty`, `qty_available`, `type`, `price`, `sale`, `is_ppn`, `active`, `created_at`, `updated_at`, `deleted_at`, `deleted_by_id` FROM `product` ORDER BY `id` DESC"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var products []domain.Product
	var createdAt, updatedAt, deletedAt sql.NullTime
	for rows.Next() {
		product := domain.Product{}
		err := rows.Scan(
			&product.Id,
			&product.ProductCategoryId,
			&product.SkuId,
			&product.Name,
			&product.Qty,
			&product.QtyAvailable,
			&product.Type,
			&product.Price,
			&product.Sale,
			&product.IsPpn,
			&product.Active,
			&createdAt,
			&updatedAt,
			&deletedAt,
			&product.DeletedById,
		)
		helper.PanicIfError(err)

		product.CreatedAt = createdAt.Time
		product.UpdatedAt = updatedAt.Time
		product.DeletedAt = deletedAt.Time

		products = append(products, product)
	}
	return products
}

// FindById implements ProductRepository.
func (p *ProductRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, productId int) (domain.Product, error) {
	SQL := "SELECT `id`, `product_category_id`, `sku_id`, `name`, `qty`, `qty_available`, `type`, `price`, `sale`, `is_ppn`, `active`, `created_at`, `updated_at`, `deleted_at`, `deleted_by_id` FROM `product` WHERE `id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, SQL, productId)
	if err != nil {
		return domain.Product{}, err
	}
	defer rows.Close()

	product := domain.Product{}
	var createdAt, updatedAt, deletedAt sql.NullTime

	if !rows.Next() {
		return product, errors.New("produk tidak ditemukan")
	}

	err = rows.Scan(
		&product.Id,
		&product.ProductCategoryId,
		&product.SkuId,
		&product.Name,
		&product.Qty,
		&product.QtyAvailable,
		&product.Type,
		&product.Price,
		&product.Sale,
		&product.IsPpn,
		&product.Active,
		&createdAt,
		&updatedAt,
		&deletedAt,
		&product.DeletedById,
	)

	if err != nil {
		return product, err
	}

	product.CreatedAt = createdAt.Time
	product.UpdatedAt = updatedAt.Time
	product.DeletedAt = deletedAt.Time

	return product, nil
}

// FindByName implements ProductRepository.
func (p *ProductRepositoryImpl) FindByName(ctx context.Context, tx *sql.Tx, productName string) (domain.Product, error) {
	SQL := "SELECT `id`, `product_category_id`, `sku_id`, `name`, `qty`, `qty_available`, `type`, `price`, `sale`, `is_ppn`, `active`, `created_at`, `updated_at`, `deleted_at`, `deleted_by_id` FROM `product` WHERE `name` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, SQL, productName)

	if err != nil {
		return domain.Product{}, err
	}

	defer rows.Close()

	product := domain.Product{}
	var createdAt, updatedAt, deletedAt sql.NullTime

	if !rows.Next() {
		return product, errors.New("produk tidak ditemukan")
	}

	err = rows.Scan(
		&product.Id,
		&product.ProductCategoryId,
		&product.SkuId,
		&product.Name,
		&product.Qty,
		&product.QtyAvailable,
		&product.Type,
		&product.Price,
		&product.Sale,
		&product.IsPpn,
		&product.Active,
		&createdAt,
		&updatedAt,
		&deletedAt,
		&product.DeletedById,
	)

	if err != nil {
		return product, err
	}

	product.CreatedAt = createdAt.Time
	product.UpdatedAt = updatedAt.Time
	product.DeletedAt = deletedAt.Time

	return product, nil
}

// Save implements ProductRepository.
func (p *ProductRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	SQL := "INSERT INTO product(`product_category_id`, `sku_id`, `name`, `qty`, `qty_available`, `type`, `price`, `sale`, `is_ppn`, `active`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(
		ctx,
		SQL,
		product.ProductCategoryId,
		product.SkuId,
		product.Name,
		product.Qty,
		product.QtyAvailable,
		product.Type,
		product.Price,
		product.Sale,
		product.IsPpn,
		product.Active,
	)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	product.Id = int(id)
	product.CreatedAt = time.Now()
	return product
}

// Update implements ProductRepository.
func (p *ProductRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	var deletedAt sql.NullTime
	SQL := "UPDATE `product` SET `product_category_id` = ? , `sku_id` = ?, `name` = ?, `qty` = ? , `qty_available` = ?, `type` = ?, `price` = ?, `sale` = ?, `is_ppn` = ?, `active` = ?, `deleted_by_id` = ?, `deleted_at` = ? WHERE `id` = ?"
	_, err := tx.ExecContext(
		ctx,
		SQL,
		product.ProductCategoryId,
		product.SkuId,
		product.Name,
		product.Qty,
		product.QtyAvailable,
		product.Type,
		product.Price,
		product.Sale,
		product.IsPpn,
		product.Active,
		product.DeletedById,
		deletedAt,
		product.Id,
	)
	helper.PanicIfError(err)
	return product
}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}
