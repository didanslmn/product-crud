package router

import (
	"golang-crud/handler"
	"golang-crud/repository"
	"golang-crud/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterProductRoutes(r *gin.RouterGroup, db *gorm.DB) {
	// inisilaisasi repository, service, dan handler
	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	product := r.Group("/products")
	{
		product.POST("/", productHandler.CreateProduct)
		product.GET("/", productHandler.GetAllProducts)
		product.GET("/newest", productHandler.GetAllProductsOrderByCreatedAtDesc)
		product.GET("/:id", productHandler.GetProduct)
		product.PUT("/:id", productHandler.UpdateProduct)
		product.DELETE("/:id", productHandler.DeleteProduct)
	}
}
