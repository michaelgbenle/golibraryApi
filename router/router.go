package router

import (
	"github.com/gin-gonic/gin"
	"golibraryApi/handler"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/books", handler.getBooks)
	router.POST("/addbook", handler.addBook)
	router.GET("/getbook/:id", handler.getBookById)
	router.PATCH("/checkout", checkOutBook)
	router.PATCH("/return", returnBook)
}
