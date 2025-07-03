package main

import (
	"github.com/gin-gonic/gin"
	"github.com/uricampos/go-api/controller"
	"github.com/uricampos/go-api/db"
	"github.com/uricampos/go-api/repository"
	"github.com/uricampos/go-api/usecase"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	// Camada de repository
	ProductRepository := repository.NewProductRepository(dbConnection)
	// Camada de usecase
	ProductUsecase := usecase.NewProductUsecase(ProductRepository)
	// Camada de controllers
	ProductController := controller.NewProductController(ProductUsecase)

	server.GET("/products", ProductController.GetProducts)
	server.POST("/product", ProductController.CreateProduct)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"messagfe": "pong",
		})
	})

	server.Run(":8080")
}
