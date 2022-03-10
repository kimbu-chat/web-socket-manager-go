package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/macmaczhl/golightweb/internal/forms"
	"github.com/macmaczhl/golightweb/internal/models"
	"github.com/macmaczhl/golightweb/internal/repositories"
)

type Books struct {
	repo *repositories.BooksRepository
}

func NewBooks() *Books {
	return &Books{repositories.NewBooksRepository()}
}

func (h *Books) Get(c *gin.Context) {
	id := c.Param("id")
	intId, _ := strconv.Atoi(id)
	book, err := h.repo.Get(intId)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, book)
}

func (h *Books) Create(c *gin.Context) {
	form := forms.CraeteBook{}
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var book *models.Book
	var err error
	if book, err = h.repo.Create(&models.Book{Name: form.Name, Description: form.Description}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, *book)
}
