package controller

import (
	"REST_API_GO/models"
	"REST_API_GO/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	services *services.ProductService
}

func NewProductController(services *services.ProductService) *ProductController {
	return &ProductController{services: services}
}

// ✅ Update GetAllProducts dengan pagination
func (c *ProductController) GetAllProducts(ctx *gin.Context) {
	// Ambil nilai page & limit dari query parameter
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	// Panggil service dengan pagination
	products, page, limit, totalPages, err := c.services.GetAllProducts(page, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Response dengan pagination info
	ctx.JSON(http.StatusOK, gin.H{
		"data":       products,
		"page":       page,
		"limit":      limit,
		"total_page": totalPages,
	})
}

// ✅ Method CreateProduct tetap sama
func (c *ProductController) CreateProduct(ctx *gin.Context) {
	var product models.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.services.CreateProduct(&product); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, product)
}
