package services

import (
	"context"
	"errors"
	"time"

	"github.com/mohnaofal/go-clean-architecture/app/models"
	"github.com/mohnaofal/go-clean-architecture/app/repository"
	"gorm.io/gorm"
)

type productsCtx struct {
	productsRepository repository.ProductRepository
}

type Products interface {
	Create(ctx context.Context, form *models.Product) (*models.Product, error)
	Update(ctx context.Context, form *models.Product) (*models.Product, error)
	View(ctx context.Context) ([]models.Product, error)
	Detail(ctx context.Context, productID int) (*models.Product, error)
	Delete(ctx context.Context, productID int) error
}

func NewProducts(productsRepository repository.ProductRepository) Products {
	return &productsCtx{productsRepository: productsRepository}
}

func (c *productsCtx) Create(ctx context.Context, form *models.Product) (*models.Product, error) {
	data, err := c.productsRepository.Upsert(ctx, form)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *productsCtx) Update(ctx context.Context, form *models.Product) (*models.Product, error) {
	if form.ID < 1 {
		return nil, errors.New("ID cannot be empty")
	}

	data, err := c.productsRepository.Get(ctx, &models.Product{Model: gorm.Model{ID: form.ID}})
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, errors.New("product not found")
	}

	data.ProductCode = form.ProductCode
	data.ProductName = form.ProductName
	data.ProductDescription = form.ProductDescription
	data.ProductQty = form.ProductQty
	data.UpdatedAt = time.Now()

	data, err = c.productsRepository.Upsert(ctx, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *productsCtx) View(ctx context.Context) ([]models.Product, error) {
	data, err := c.productsRepository.Select(ctx, &models.Product{})
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *productsCtx) Detail(ctx context.Context, productID int) (*models.Product, error) {
	if productID < 1 {
		return nil, errors.New("ID cannot be empty")
	}

	data, err := c.productsRepository.Get(ctx, &models.Product{Model: gorm.Model{ID: uint(productID)}})
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *productsCtx) Delete(ctx context.Context, productID int) error {
	if productID < 1 {
		return errors.New("ID cannot be empty")
	}

	return c.productsRepository.Delete(ctx, &models.Product{Model: gorm.Model{ID: uint(productID)}})
}
