package router

import "github.com/gin-gonic/gin"

router.GET("/books", getBooks)
router.POST("/addbook", addBook)
router.GET("/getbook/:id", getBookById)
router.PATCH("/checkout", checkOutBook)
router.PATCH("/return", returnBook)
err := router.Run("localhost:8080")
if err != nil {
return
}

func SetupRouter() *gin.Engine{
	router := gin.Default()
}