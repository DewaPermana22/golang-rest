package controller

import (
	"REST_API_GO/models"
	"REST_API_GO/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	services *services.CategoryService
}

func NewCategoryController(services *services.CategoryService) *CategoryController {
	return &CategoryController{services: services}
}

func (c *CategoryController) GetAllCategory(ctx *gin.Context)  {
	products, err := c.services.GetAllCategory()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, products)
}
func (c *CategoryController) CreateCategory(ctx *gin.Context)  {
	var categories models.Category
	if err := ctx.ShouldBindJSON(&categories); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.services.CreateCategory(&categories); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, categories)
}