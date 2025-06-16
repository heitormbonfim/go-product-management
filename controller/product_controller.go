package controller

import (
	"net/http"
	"product-management/model"
	"product-management/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productUsecase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) productController {
	return productController{
		productUsecase: usecase,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {
	products, err := p.productUsecase.GetProducts()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, products)
}

func (p *productController) AddProduct(ctx *gin.Context) {
	var product model.Product
	err := ctx.BindJSON(&product)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedProduct, err := p.productUsecase.AddProduct(product)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusCreated, insertedProduct)
}

func (p *productController) GetProductById(ctx *gin.Context) {
	id := ctx.Param("productId")

	if id == "" {
		response := model.Response{
			Message: "Product id cannot be null",
		}

		ctx.JSON(http.StatusBadRequest, response)

		return
	}

	productId, err := strconv.Atoi(id)

	if err != nil {
		response := model.Response{
			Message: "Product id needs to be a integer number",
		}

		ctx.JSON(http.StatusBadRequest, response)

		return
	}

	product, err := p.productUsecase.GetProductById(productId)

	if err != nil && product == nil {
		response := model.Response{
			Message: "Product not found",
		}

		ctx.JSON(http.StatusNotFound, response)

		return
	}

	ctx.JSON(http.StatusOK, product)
}
