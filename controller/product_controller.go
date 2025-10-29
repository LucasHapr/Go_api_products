package controller

import (
	"database/sql"
	"errors"
	"go-api/model"
	"go-api/usecase"
	"net/http"
	"strconv"

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

func (p *ProductController) CreateProduct(ctx *gin.Context) {
	var product model.Product
	err := ctx.BindJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdProduct, err := p.ProductUsecase.CreateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, createdProduct)
}

func (p *ProductController) GetProductByID(ctx *gin.Context) {
	var product *model.Product
	idParam := ctx.Param("id")
	if idParam == "" {
		response := model.Response{Message: "Product ID is required"}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productID, err := strconv.Atoi(idParam)
	if err != nil {
		response := model.Response{Message: "Invalid Product ID"}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product, err = p.ProductUsecase.GetProductByID(productID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if product == nil {
		response := model.Response{Message: "Product not found"}
		ctx.JSON(http.StatusNotFound, response)
		return
	}
	ctx.JSON(http.StatusOK, product)
}

func (p *ProductController) DeleteProduct(ctx *gin.Context) {
	idParam := ctx.Param("id")
	if idParam == "" {
		response := model.Response{Message: "Product ID is required"}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productID, err := strconv.Atoi(idParam)
	if err != nil {
		response := model.Response{Message: "Invalid Product ID"}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err = p.ProductUsecase.DeleteProduct(productID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response := model.Response{Message: "Product deleted successfully"}
	ctx.JSON(http.StatusOK, response)
}

func (p *ProductController) UpdateProduct(ctx *gin.Context) {
	var product model.Product
	idParam := ctx.Param("id")
	if idParam == "" {
		response := model.Response{Message: "Product ID is required"}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productID, err := strconv.Atoi(idParam)
	if err != nil {
		response := model.Response{Message: "Invalid Product ID"}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	product.ID = productID
	err = ctx.BindJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = p.ProductUsecase.UpdateProduct(product)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response := model.Response{Message: "Product not found"}
			ctx.JSON(http.StatusNotFound, response)
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response := model.Response{Message: "Product updated successfully"}
	ctx.JSON(http.StatusOK, response)
}
