package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()
	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	//repository
	ProductRepository := repository.NewProductRepository(dbConnection)

	//usecase
	ProductUseCase := usecase.NewProductUsecase(ProductRepository)

	//controller
	ProductController := controller.NewProductController(ProductUseCase)

	//routes
	server.GET("/products", ProductController.GetProducts)
	server.GET("/products/:id", ProductController.GetProductByID)

	server.POST("/products", ProductController.CreateProduct)

	server.PUT("/products/:id", ProductController.UpdateProduct)

	server.DELETE("/products/:id", ProductController.DeleteProduct)

	server.Run(":8080")

}
