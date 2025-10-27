package controller

import (
	"go-api/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductUsecase *usecase.ProductUsecase
}

func NewProductController(usecase *usecase.ProductUsecase) *ProductController {
	return &ProductController{
		ProductUsecase: usecase,
	}
}

func (pc *ProductController) GetProducts(ctx *gin.Context) {
	products, err := pc.ProductUsecase.GetAllProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, products)
}
