package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": books})
}
func addBooks(c *gin.Context) {
	var newBook book
	err := c.ShouldBindJSON(newBook)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "unable to bind json"})
		c.Abort()
		return
	}
	books = append(books, newBook)

}
func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "Rules of life", Author: "Jordan Peterson", Quantity: 12},
	{ID: "2", Title: "Rules of love", Author: "Kate Peterson", Quantity: 10},
	{ID: "3", Title: "Rules of wealth", Author: "Jordan Peterson", Quantity: 12},
	{ID: "4", Title: "Rules of health", Author: "Kate Peterson", Quantity: 11},
}
