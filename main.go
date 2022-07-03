package main

import (
	"github.com/gin-gonic/gin"
)

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

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/addbook", addBook)
	router.GET("/getbook/:id", getBookById)
	router.PATCH("/checkout", checkOutBook)
	router.PATCH("/return", returnBook)
	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}
