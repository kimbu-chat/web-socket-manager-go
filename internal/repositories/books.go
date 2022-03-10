package repositories

import (
	"gorm.io/gorm"

	"github.com/kimbu-chat/web-socket-manager-go/internal/config/db"
	"github.com/kimbu-chat/web-socket-manager-go/internal/models"
)

type BooksRepository struct {
	db *gorm.DB
}

func NewBooksRepository() *BooksRepository {
	return &BooksRepository{db.Connection()}
}

func (r *BooksRepository) Get(id int) (*models.Book, error) {
	book := models.Book{Id: id}
	return &book, r.db.First(&book).Error
}

func (r *BooksRepository) Create(book *models.Book) (*models.Book, error) {
	err := r.db.Create((book)).Error
	return book, err
}
