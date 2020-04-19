package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rahmanfadhil/gin-bookstore/controllers"
	"github.com/rahmanfadhil/gin-bookstore/models"
	_ "net/http"
	//"Book/controllers"
	//"Book/models"
)

func main()  {
	r := gin.Default()

	db := models.SetupModels()

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.GET("/", controllers.FindBooks)
	r.POST("/book", controllers.CreateBook)
	r.GET("/book/:id", controllers.FindBook)
	r.PATCH("/book/:id", controllers.UpdateBook)
	r.DELETE("/book/:id", controllers.DeleteBook)

	r.Run()
}