package repository

import (
	"context"

	"github.com/mohnaofal/go-clean-architecture/app/models"
	"gorm.io/gorm"
)

type productRepositoryCtx struct {
	db *gorm.DB
}

type ProductRepository interface {
	Upsert(ctx context.Context, data *models.Product) (*models.Product, error)
	Delete(ctx context.Context, data *models.Product) error
	Select(ctx context.Context, data *models.Product) ([]models.Product, error)
	Get(ctx context.Context, data *models.Product) (*models.Product, error)
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepositoryCtx{db: db}
}

func (c *productRepositoryCtx) Upsert(ctx context.Context, data *models.Product) (*models.Product, error) {
	if err := c.db.Save(data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (c *productRepositoryCtx) Select(ctx context.Context, data *models.Product) ([]models.Product, error) {
	rows := make([]models.Product, 0)
	if err := c.db.Find(&rows).Error; err != nil {
		return rows, err
	}
	return rows, nil
}

func (c *productRepositoryCtx) Get(ctx context.Context, data *models.Product) (*models.Product, error) {
	if err := c.db.First(data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (c *productRepositoryCtx) Delete(ctx context.Context, data *models.Product) error {
	if err := c.db.Delete(data).Error; err != nil {
		return err
	}
	return nil
}
