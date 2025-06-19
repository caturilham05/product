package service

import (
	"caturilham05/product/exception"
	"caturilham05/product/model/domain"
	"context"
	"database/sql"
	"errors"
	"log"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestMain(m *testing.M) {
	log.Println("Starting Product Service Tests")
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	log.Println("Environment variables loaded successfully")
	os.Exit(m.Run())
}

type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) FindById(ctx context.Context, tx *sql.Tx, productId int) (domain.Product, error) {
	args := m.Called(ctx, tx, productId)
	if args.Get(0) == nil {
		return domain.Product{}, args.Error(1)
	}
	return args.Get(0).(domain.Product), args.Error(1)
}

func (m *MockProductRepository) Delete(ctx context.Context, tx *sql.Tx, productId int, userId int) {
	m.Called(ctx, tx, productId, userId)
}

func (m *MockProductRepository) FindAll(ctx context.Context, tx *sql.Tx) []domain.Product {
	args := m.Called()
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).([]domain.Product)
}

func (m *MockProductRepository) FindByName(ctx context.Context, tx *sql.Tx, productName string) (domain.Product, error) {
	args := m.Called(ctx, tx, productName)
	if args.Get(0) == nil {
		return domain.Product{}, args.Error(1)
	}
	return args.Get(0).(domain.Product), args.Error(1)
}

func (m *MockProductRepository) Save(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	args := m.Called(ctx, tx, product)
	if args.Get(0) == nil {
		return domain.Product{}
	}
	return args.Get(0).(domain.Product)
}

func (m *MockProductRepository) Update(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	args := m.Called(ctx, tx, product)
	if args.Get(0) == nil {
		return domain.Product{}
	}
	return args.Get(0).(domain.Product)
}

func TestServiceImpl_FindById_Success(t *testing.T) {
	db, mockDB, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer db.Close()
	mockRepo := new(MockProductRepository)
	productService := NewProductService(mockRepo, db, nil)

	expectedDomainProduct := domain.Product{
		Id: 1,
	}

	mockDB.ExpectBegin()
	mockRepo.On("FindById", mock.Anything, mock.Anything, 1).Return(expectedDomainProduct, nil)
	mockDB.ExpectCommit()
	productResponse := productService.FindById(context.Background(), 1)
	assert.NoError(t, nil, "Expected no error")
	assert.NotNil(t, productResponse, "Expected product response to be not nil")
	assert.Equal(t, expectedDomainProduct.Id, productResponse.Id, "Expected product ID to match")
	assert.Equal(t, expectedDomainProduct.Name, productResponse.Name, "Expected product name to match")

	mockRepo.AssertExpectations(t)
}

func TestServiceImpl_FindById_NotFound(t *testing.T) {
	db, mockDB, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer db.Close()
	mockRepo := new(MockProductRepository)
	productService := NewProductService(mockRepo, db, nil)

	mockDB.ExpectBegin()
	mockRepo.On("FindById", mock.Anything, mock.Anything, 100).Return(domain.Product{}, exception.NewNotFoundError("product tidak ditemukan"))
	mockDB.ExpectRollback()

	assert.Panics(t, func() {
		productService.FindById(context.Background(), 100)
	}, "Expected panic due to not found error")

	defer func() {
		if r := recover(); r != nil {
			err, ok := r.(exception.NotFoundError)
			assert.True(t, ok, "panic seharusnya bertipe NotFoundError")
			assert.Equal(t, "produk tidak ditemukan", err.Error)
		}
	}()

	mockRepo.AssertExpectations(t)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

func TestServiceImpl_FindById_InternalError(t *testing.T) {
	db, mockDB, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer db.Close()
	mockRepo := new(MockProductRepository)
	ProductService := NewProductService(mockRepo, db, nil)

	mockDB.ExpectBegin()
	expectedError := errors.New("internal server error")
	mockRepo.On("FindById", mock.Anything, mock.Anything, 2).Return(domain.Product{}, expectedError)
	mockDB.ExpectRollback()

	assert.PanicsWithValue(t, expectedError, func() {
		ProductService.FindById(context.Background(), 2)
	}, "service seharusnya panic dengan error internal server")
	mockRepo.AssertExpectations(t)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}
