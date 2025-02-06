package db

import (
	"test-exam-forviz/internal/models"
)

type BookRepository interface {
	Create(book models.BookRepository) error
	Update(book models.BookRepository) error
	Delete(id int) error
	FindByID(id int) (models.BookRepository, error)
	FindAll(title, author, category, sortName, sortType string) ([]models.BookRepository, error)
	BorrowBook(id, count int) error
	ReturnBook(id int) error
}
