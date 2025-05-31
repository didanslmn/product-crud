package main

import (
	"golang-crud/config"
	router_product "golang-crud/router"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("failed to connect databse:%v", err)
	}
	router := gin.Default()

	api := router.Group("/api")

	router_product.RegisterProductRoutes(api, db)

	router.Run(":8080")

}
