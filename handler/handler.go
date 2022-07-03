package handler

import (
	"github.com/gin-gonic/gin"
	"golibraryApi/database"
	"net/http"
)

type Handler struct {
	DB database.DB
}

func bookById(id string) (*book, error) {
	for i, v := range books {
		if v.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")
}
func (h *Handler) GetBooks(c *gin.Context) {
	books, err := h.DB.GetAllBooks()
	if err != nil {
		c.IndentedJSON(http.StatusOK, gin.H{"message": books})
		return
	}

}
func (h *Handler) GetBookById(c *gin.Context) {
	id := c.Param("id")
	sBook, err := bookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "book not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": sBook})
}
func (h *Handler) AddBook(c *gin.Context) {
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
func (h *Handler) CheckOutBook(c *gin.Context) {
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
