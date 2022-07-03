package router

import (
	"github.com/gin-gonic/gin"
	"golibraryApi/handler"
	"os"
)

func SetupRouter(h *handler.Handler) (*gin.Engine, string) {
	router := gin.Default()
	router.GET("/books", handler.GetBooks)
	router.POST("/addbook", handler.AddBook)
	router.GET("/getbook/:id", handler.GetBookById)
	router.PATCH("/checkout", handler.CheckOutBook)
	router.PATCH("/return", handler.ReturnBook)
	port := os.Getenv("PORT")

	return router, port
}
