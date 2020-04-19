package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/rahmanfadhil/gin-bookstore/models"
	"net/http"
)

type CreateBookInput struct {
	Title string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

type UpdateBook struct {
	Title string `json:"title"`
	Author string `json:"author"`
}

func FindBooks(c *gin.Context)  {
	db := c.MustGet("db").(*gorm.DB)

	var books []models.Book

	db.Find(&books)
	c.JSON(http.StatusOK, gin.H{"data":books})
}

func createBook(c *gin.Context)  {
	db := c.MustGet("db").(*gorm.DB)

	//validate input
	var input CreateBookInput
	if err := c.ShouldBindJSON(&input);
	err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//create a book
	book := models.Book{
		Title: input.Title,
		Author: input.Author,
	}

	db.Create(&book)
	c.JSON(http.StatusOK, gin.H{"status":"succss", "data":book})
}

func findBook(c *gin.Context)  {
	db := c.MustGet("db").(*gorm.DB)

	var book models.Book
	if err := db.Where("id=?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "No Record Found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func updateBook(c *gin.Context)  {
	db := c.MustGet("db").(*gorm.DB)

	var book models.Book
	if error := db.Where("id=?", c.Param("id")).First(&book).Error; error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data":"Record Not Found"})
		return
	}

	var input UpdateBook
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Invalid Input Date"})
		return
	}

	db.Model(&book).Update(input)
	c.JSON(http.StatusOK, gin.H{"status": "Record Updated Successfully", "data":book})
}

func deleteBook(c gin.Context)  {
	db := c.MustGet("db").(*gorm.DB)

	var book models.Book
	if err := db.Where("id=?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "No Record Found"})
		return
	}

	db.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"data": "Record Delete Successfully"})
}
