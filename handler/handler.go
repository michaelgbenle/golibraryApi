package handler

import (
	"github.com/gin-gonic/gin"
	"golibraryApi/database"
	"golibraryApi/models"
	"log"
	"net/http"
)

type Handler struct {
	DB database.DB
}

func (h *Handler) GetBooks(c *gin.Context) {
	books, err := h.DB.GetAllBooks()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": "error fetching books"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": books})
}

func (h *Handler) GetBookById(c *gin.Context) {
	id := c.Param("id")
	book, err := h.DB.BookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "book not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": book})
}
func (h *Handler) AddBook(c *gin.Context) {
	var newBook models.Book
	err := c.ShouldBindJSON(&newBook)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "unable to bind json"})
		return
	}

	lerr := h.DB.AddNewBook(newBook)

	if lerr != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "unable to add book"})
		return
	}
	c.IndentedJSON(http.StatusCreated, gin.H{
		"message": "book successfully added",
		"book":    newBook,
	})
}

func (h *Handler) CheckOutBook(c *gin.Context) {
	//info := &models.CheckOutBook{}
	//c.ShouldBindJSON(info)
	//book, err := h.DB.BookById(info.Id)
	id := c.Query("id")
	copies := c.Query("copies")

	book, err := h.DB.BookById(id)
	log.Println(book)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "book not found"})
		return
	}
	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "book not available"})
		return
	}
	log.Println(id)
	log.Println(copies)
	newBook, berr := h.DB.Checkout(id, copies)
	if berr != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": " could not checkout book"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "book checked out successfully ",
		"updated": newBook,
	})
}

func (h *Handler) ReturnBook(c *gin.Context) {
	id := c.Query("id")
	copies := c.Query("copies")
	//book, err := h.DB.BookById(id)
	//if err != nil {
	//	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "book not found"})
	//	return
	//}
	newBook, berr := h.DB.Checkin(id, copies)
	if berr != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": " could not update"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "book successfully returned",
		"updated": newBook,
	})
}
