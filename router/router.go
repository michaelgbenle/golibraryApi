package router

import (
	"github.com/gin-gonic/gin"
	"golibraryApi/handler"
	"os"
)

func SetupRouter(h *handler.Handler) (*gin.Engine, string) {
	router := gin.Default()
	router.GET("/books", h.GetBooks)
	router.POST("/addbook", h.AddBook)
	router.GET("/getbook/:id", h.GetBookById)
	router.PATCH("/checkout", h.CheckOutBook)
	router.PATCH("/return", h.ReturnBook)
	router.DELETE("/delete", h.DeleteBook)

	port := os.Getenv("PORT")

	return router, port
}
