package services

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/mohnaofal/go-clean-architecture/app/models"
	"gorm.io/gorm"
)

var products = make([]models.Product, 0)

type productsCtx struct {
}

type Products interface {
	Create(ctx context.Context, form *models.Product) (*models.Product, error)
	Update(ctx context.Context, form *models.Product) (*models.Product, error)
	View(ctx context.Context) ([]models.Product, error)
	Detail(ctx context.Context, productCode string) (*models.Product, error)
	Delete(ctx context.Context, productCode string) error
}

func NewProducts() Products {
	return &productsCtx{}
}

func (c *productsCtx) Create(ctx context.Context, form *models.Product) (*models.Product, error) {
	for _, val := range products {
		if strings.EqualFold(val.ProductCode, form.ProductCode) {
			return nil, errors.New("product already exist")
		}
	}

	data := &models.Product{
		Model: gorm.Model{
			ID:        uint(time.Now().Unix()),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		ProductCode:        form.ProductCode,
		ProductName:        form.ProductName,
		ProductDescription: form.ProductDescription,
		ProductQty:         form.ProductQty,
	}

	products = append(products, *data)

	return data, nil
}
func (c *productsCtx) Update(ctx context.Context, form *models.Product) (*models.Product, error) {
	isExist := false
	data := new(models.Product)
	for i, val := range products {
		if strings.EqualFold(val.ProductCode, form.ProductCode) {
			isExist = true
			data = &models.Product{
				Model: gorm.Model{
					ID:        val.ID,
					CreatedAt: val.CreatedAt,
					UpdatedAt: time.Now(),
				},
				ProductCode:        form.ProductCode,
				ProductName:        form.ProductName,
				ProductDescription: form.ProductDescription,
				ProductQty:         form.ProductQty,
			}

			products[i] = *data
			break
		}
	}

	if !isExist {
		return nil, errors.New("product not found")
	}

	return data, nil
}

func (c *productsCtx) View(ctx context.Context) ([]models.Product, error) {
	return products, nil
}

func (c *productsCtx) Detail(ctx context.Context, productCode string) (*models.Product, error) {
	isExist := false
	data := new(models.Product)
	for _, val := range products {
		if strings.EqualFold(val.ProductCode, productCode) {
			isExist = true
			data = &val
			break
		}
	}

	if !isExist {
		return nil, errors.New("product not found")
	}

	return data, nil
}

func (c *productsCtx) Delete(ctx context.Context, productCode string) error {
	isExist := false
	for i, val := range products {
		if strings.EqualFold(val.ProductCode, productCode) {
			isExist = true
			products = append(products[:i], products[i+1:]...)
			break
		}
	}

	if !isExist {
		return errors.New("product not found")
	}

	return nil
}
