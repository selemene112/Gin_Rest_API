package main

import (
	"API_PSQL/controller"
	"API_PSQL/models"

	"github.com/gin-gonic/gin"
)



func main() {
	r := gin.Default()
	models.ConnectDatabase()
	
	r.POST("/Products", controller.Create)
	r.GET("/Products", controller.Show)
	r.GET("/Products:id", controller.Getbyid)
	r.PUT("/Products:id", controller.Update)
	r.DELETE("/Products", controller.Delete)



	// controller.Create() 	
	// controller.Show()			
	r.Run()
}