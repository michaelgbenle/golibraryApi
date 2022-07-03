package router

import (
	"github.com/gin-gonic/gin"
	"golibraryApi/handler"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/books", handler.GetBooks)
	router.POST("/addbook", handler.AddBook)
	router.GET("/getbook/:id", handler.AetBookById)
	router.PATCH("/checkout", handler.AheckOutBook)
	router.PATCH("/return", handler.ReturnBook)

	return router
}
