package main

import (
	"test1/database"
	"test1/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/product", handlers.GetProduct)
		api.GET("/product/:id", handlers.GetProductID)
		api.POST("/product", handlers.CreateProduct)
		api.PUT("/product/:id", handlers.UpdateProduct)
		api.DELETE("/product/:id", handlers.DeleteProduct)

		api.GET("/measure", handlers.GetMeasure)
		api.GET("/measure/:id", handlers.GetMeasureID)
		api.POST("/measure", handlers.CreateMeasure)
		api.PUT("/measure/:id", handlers.UpdateMeasure)
		api.DELETE("/measure/:id", handlers.DeleteMeasure)
	}

	r.Run(":8080")
}
