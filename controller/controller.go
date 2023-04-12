package controller

import (
	"net/http"

	"project/database"
	"project/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetBooks(c *gin.Context) {

	db := database.GetDB()

	var books []models.Book

	result := db.Find(&books)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": result.Error.Error()})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"Message": "No books found"})
		return
	}

	c.JSON(http.StatusOK, books)
}

func GetBookByID(c *gin.Context) {

	db := database.GetDB()

	var book models.Book

	id := c.Param("id")
	result := db.First(&book, id)

	if result.Error != nil {

		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"Error": "Book Not Found!!"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"Error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

func CreateBook(c *gin.Context) {

	db := database.GetDB()

	var book models.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	result := db.Create(&book)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Successfully added new book", "book": book})
}

func UpdateBook(c *gin.Context) {

	db := database.GetDB()

	var book models.Book

	id := c.Param("id")
	result := db.First(&book, id)

	if result.Error != nil {

		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result = db.Save(&book)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book updated successfully", "book": book})
}

func DeleteBook(c *gin.Context) {

	db := database.GetDB()

	id := c.Param("id")
	result := db.Delete(&models.Book{}, id)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
