package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func bookById(id string) (*book, error) {
	for i, v := range books {
		if v.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")
}
func GetBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": books})
}
func GetBookById(c *gin.Context) {
	id := c.Param("id")
	sBook, err := bookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "book not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": sBook})
}
func AddBook(c *gin.Context) {
	var newBook book
	err := c.ShouldBindJSON(&newBook)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "unable to bind json"})
		c.Abort()
		return
	}
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, gin.H{"message": newBook})
}
func CheckOutBook(c *gin.Context) {
	id := c.Query("id")
	sBook, err := bookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "book not found"})
		return
	}
	if sBook.Quantity <= 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "book not available"})
		return
	}
	sBook.Quantity -= 1
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "book successfully returned",
		"updated": sBook,
	})
}

func returnBook(c *gin.Context) {
	id := c.Query("id")
	sBook, err := bookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "book not found"})
		return
	}
	sBook.Quantity += 1
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "book successfully returned",
		"updated": sBook,
	})
}
