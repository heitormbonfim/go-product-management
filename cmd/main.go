package main

import (
	"product-management/controller"
	"product-management/db"
	"product-management/repository"
	"product-management/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()

	if err != nil {
		panic(err)
	}

	ProductRepository := repository.NewProductRepository(dbConnection)

	ProductUsecase := usecase.NewProductUsecase(ProductRepository)

	ProductController := controller.NewProductController(ProductUsecase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.POST("/product", ProductController.AddProduct)
	server.GET("/products", ProductController.GetProducts)
	server.GET("/products/:productId", ProductController.GetProductById)

	server.Run(":8000")
}
