package service

import (
	"context"
	"fmt"
	"golang-crud/dto"
	"golang-crud/model"
	"golang-crud/repository"
	"log"
)

type ProductService interface {
	Create(ctx context.Context, req *dto.CreateProductRequest) (*dto.ProductResponse, error)
	GetByID(ctx context.Context, id uint) (*dto.ProductResponse, error)
	GetAll(ctx context.Context) ([]dto.ProductResponse, error)
	GetAllOrderByCreatedAtDesc(ctx context.Context) ([]dto.ProductResponse, error)
	Update(ctx context.Context, id uint, req *dto.UpdateProductRequest) (*dto.ProductResponse, error)
	Delete(ctx context.Context, id uint) error
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) Create(ctx context.Context, req *dto.CreateProductRequest) (*dto.ProductResponse, error) {
	product := &model.Product{
		Nama:      req.Nama,
		Deskripsi: req.Deskripsi,
		Harga:     req.Harga,
		Kategori:  req.Kategori,
	}

	if err := s.repo.Create(ctx, product); err != nil {
		log.Printf("Failed to create product: %v", err)
		return nil, fmt.Errorf("failed to create product: %w", err)
	}

	log.Printf("Product created: %s", product.Nama)

	return &dto.ProductResponse{
		ID:        int(product.ID),
		Nama:      product.Nama,
		Deskripsi: product.Deskripsi,
		Harga:     product.Harga,
		Kategori:  product.Kategori,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}, nil
}

func (s *productService) GetByID(ctx context.Context, id uint) (*dto.ProductResponse, error) {
	product, err := s.repo.GetByID(ctx, id)
	if err != nil {
		log.Printf("Product not found (id: %d): %v", id, err)
		return nil, fmt.Errorf("product not found: %w", err)
	}

	log.Printf("Fetched product: %s", product.Nama)

	return &dto.ProductResponse{
		ID:        int(product.ID),
		Nama:      product.Nama,
		Deskripsi: product.Deskripsi,
		Harga:     product.Harga,
		Kategori:  product.Kategori,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}, nil
}

func (s *productService) GetAll(ctx context.Context) ([]dto.ProductResponse, error) {
	products, err := s.repo.GetAll(ctx)
	if err != nil {
		log.Printf("Failed to fetch products: %v", err)
		return nil, fmt.Errorf("failed to fetch products: %w", err)
	}

	log.Println("Fetched all products")

	var responses []dto.ProductResponse
	for _, p := range products {
		responses = append(responses, dto.ProductResponse{
			ID:        int(p.ID),
			Nama:      p.Nama,
			Deskripsi: p.Deskripsi,
			Harga:     p.Harga,
			Kategori:  p.Kategori,
			CreatedAt: p.CreatedAt,
			UpdatedAt: p.UpdatedAt,
		})
	}

	return responses, nil
}
func (s *productService) GetAllOrderByCreatedAtDesc(ctx context.Context) ([]dto.ProductResponse, error) {
	products, err := s.repo.GetAllOrderByCreatedAtDesc(ctx)
	if err != nil {
		log.Printf("failed to fetch product ordered bt created_at desc: %v", err)
		return nil, fmt.Errorf("failed to fetch products ordered by created_at desc: %w", err)
	}

	var responses []dto.ProductResponse
	for _, p := range products {
		responses = append(responses, dto.ProductResponse{
			ID:        int(p.ID),
			Nama:      p.Nama,
			Deskripsi: p.Deskripsi,
			Harga:     p.Harga,
			Kategori:  p.Kategori,
			CreatedAt: p.CreatedAt,
			UpdatedAt: p.UpdatedAt,
		})
	}
	return responses, nil
}

func (s *productService) Update(ctx context.Context, id uint, req *dto.UpdateProductRequest) (*dto.ProductResponse, error) {
	product, err := s.repo.GetByID(ctx, id)
	if err != nil {
		log.Printf("Product not found (id: %d): %v", id, err)
		return nil, fmt.Errorf("product not found: %w", err)
	}

	product.Nama = req.Nama
	product.Deskripsi = req.Deskripsi
	product.Harga = req.Harga
	product.Kategori = req.Kategori

	if err := s.repo.Update(ctx, id, product); err != nil {
		log.Printf("Failed to update product (id: %d): %v", id, err)
		return nil, fmt.Errorf("failed to update product: %w", err)
	}

	log.Printf("Updated product (id: %d)", id)

	return &dto.ProductResponse{
		ID:        int(product.ID),
		Nama:      product.Nama,
		Deskripsi: product.Deskripsi,
		Harga:     product.Harga,
		Kategori:  product.Kategori,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}, nil
}

func (s *productService) Delete(ctx context.Context, id uint) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		log.Printf("Failed to delete product (id: %d): %v", id, err)
		return fmt.Errorf("failed to delete product: %w", err)
	}

	log.Printf("Deleted product (id: %d)", id)
	return nil
}
