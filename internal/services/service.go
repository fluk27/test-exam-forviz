package services

import (
	"test-exam-forviz/internal/models"
)

type BookService interface {
	CreateBook(book models.BookRequest) (models.BookResponse, error)
	UpdateBook(id int, book models.BookRequest) (models.BookResponse, error)
	DeleteBook(id int) (models.BookResponse, error)
	GetBookByID(id int) (models.BookResponse, error)
	SearchBooks(title, author, category string) (models.BookListResponse, error)
	GetMostBorrowedBooks() (models.BookListResponse, error)
	BorrowBook(id int) (models.BookResponse, error)
	ReturnBook(id int) (models.BookResponse, error)
}
