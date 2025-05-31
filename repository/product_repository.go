package repository

import (
	"context"
	"errors"
	"fmt"
	"golang-crud/model"
	"log"

	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(ctx context.Context, product *model.Product) error
	GetByID(ctx context.Context, id uint) (*model.Product, error)
	GetAll(ctx context.Context) ([]model.Product, error)
	GetAllOrderByCreatedAtDesc(ctx context.Context) ([]model.Product, error)
	Update(ctx context.Context, id uint, product *model.Product) error
	Delete(ctx context.Context, id uint) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}
func (r *productRepository) Create(ctx context.Context, product *model.Product) error {
	if err := r.db.WithContext(ctx).Create(product).Error; err != nil {
		log.Printf("create product failed (name: %s): %v", product.Nama, err)
		return fmt.Errorf("failed to create product: %w", err)
	}
	return nil
}

func (r *productRepository) GetByID(ctx context.Context, id uint) (*model.Product, error) {
	var product model.Product
	if err := r.db.WithContext(ctx).First(&product, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}
		log.Printf("Get product bt ID failed (id: %d): %v", id, err)
		return nil, fmt.Errorf("failed to get product by ID: %w", err)
	}
	return &product, nil
}
func (r *productRepository) GetAll(ctx context.Context) ([]model.Product, error) {
	var products []model.Product
	if err := r.db.WithContext(ctx).Find(&products).Error; err != nil {
		log.Printf("Get all products failed: %v", err)
		return nil, fmt.Errorf("failed to get products: %w", err)
	}
	return products, nil
}
func (r *productRepository) GetAllOrderByCreatedAtDesc(ctx context.Context) ([]model.Product, error) {
	var products []model.Product
	if err := r.db.WithContext(ctx).Order("created_at desc").Find(&products).Error; err != nil {
		log.Printf("Get all products ordered by created_at desc failed: %v", err)
		return nil, fmt.Errorf("failed to get products ordered by created_at desc: %w", err)
	}
	return products, nil
}

func (r *productRepository) Update(ctx context.Context, id uint, input *model.Product) error {
	var product model.Product
	if err := r.db.WithContext(ctx).First(&product, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("Update product failed - not found (id: %d)", id)
			return errors.New("product not found")
		}
		log.Printf("Update product failed (id: %d): %v", id, err)
		return fmt.Errorf("failed to find product by ID: %w", err)
	}

	product.Nama = input.Nama
	product.Deskripsi = input.Deskripsi
	product.Harga = input.Harga
	product.Kategori = input.Kategori

	if err := r.db.WithContext(ctx).Save(&product).Error; err != nil {
		log.Printf("Save updated product failed (id: %d): %v", id, err)
		return fmt.Errorf("failed to update product: %w", err)
	}
	return nil
}
func (r *productRepository) Delete(ctx context.Context, id uint) error {
	result := r.db.WithContext(ctx).Model(&model.Product{}).Where("id = ?", id).Delete(&model.Product{})
	if result.Error != nil {
		return fmt.Errorf("failed to soft delete product: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		log.Printf("Delete product failed - not found (id: %d)", id)
		return errors.New("product not found")
	}
	return nil

}
