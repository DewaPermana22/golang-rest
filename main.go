package main

import (
	"REST_API_GO/routes"
	"REST_API_GO/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {
	utils.Connection()

	r := gin.Default()
	r.Use(cors.Default()) // Mengizinkan semua origin untuk mengakses API
	routes.InitRoutes(r)



	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, Gin!"})
	})

	r.Run(":5000")
	//database.MigrationDB()
}
