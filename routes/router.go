package routes

import (
	"REST_API_GO/controller"
	"REST_API_GO/repositories"
	"REST_API_GO/services"
	"REST_API_GO/utils"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine)  {
	productsRepo := repositories.NewProductRepository(utils.DB)
	productServices := services.NewProductService(productsRepo)
	productController := controller.NewProductController(productServices)

	//Category
	categoriesRepo := repositories.NewCategoryRepository(utils.DB)
	categoriesServices := services.NewCategoryService(categoriesRepo)
	categoriesController := controller.NewCategoryController(categoriesServices)

	r.GET("/products", productController.GetAllProducts)
	r.POST("/products/add", productController.CreateProduct)

	r.GET("/categories", categoriesController.GetAllCategory)
	r.POST("/categories/add", categoriesController.CreateCategory)
}