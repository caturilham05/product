package repository

import (
	"caturilham05/product/model/domain"
	"context"
	"database/sql"
)

type ProductRepository interface {
	Save(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product
	Update(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product
	Delete(ctx context.Context, tx *sql.Tx, productId int, userId int)
	FindById(ctx context.Context, tx *sql.Tx, productId int) (domain.Product, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Product
	FindByName(ctx context.Context, tx *sql.Tx, productName string) (domain.Product, error)
}
