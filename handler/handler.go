package handler

import (
	"github.com/gin-gonic/gin"
	"golibraryApi/database"
	"golibraryApi/models"
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
	err = h.DB.AddNewBook(newBook)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "unable to add book"})
		return
	}
	c.IndentedJSON(http.StatusCreated, gin.H{
		"message": "book successfully added",
		"book":    newBook,
	})
}

func (h *Handler) CheckOutBook(c *gin.Context) {
	id := c.Query("id")
	copies := c.Query("copies")
	book, err := h.DB.BookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "book not found"})
		return
	}
	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "book not available"})
		return
	}
	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "book successfully returned",
		"updated": book,
	})
}

func (h *Handler) ReturnBook(c *gin.Context) {
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
