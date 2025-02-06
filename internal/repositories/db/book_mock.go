package db

import (
	"test-exam-forviz/internal/models"

	"github.com/stretchr/testify/mock"
)

type mockBookRepository struct {
	mock.Mock
}

func (mockBookRepo *mockBookRepository) Create(book models.BookRepository) error {
	args := mockBookRepo.Called()
	return args.Error(0)
}
func (mockBookRepo *mockBookRepository) Update(book models.BookRepository) error {
	args := mockBookRepo.Called()
	return args.Error(0)
}
func (mockBookRepo *mockBookRepository) Delete(id int) error {
	args := mockBookRepo.Called()
	return args.Error(0)
}
func (mockBookRepo *mockBookRepository) FindByID(id int) (models.BookRepository, error) {
	args := mockBookRepo.Called()
	return args.Get(0).(models.BookRepository), args.Error(1)
}
func (mockBookRepo *mockBookRepository) FindAll(title, author, category, sortName, sortType string) ([]models.BookRepository, error) {
	args := mockBookRepo.Called()
	return args.Get(0).([]models.BookRepository), args.Error(1)
}
func (mockBookRepo *mockBookRepository) BorrowBook(id, count int) error {
	args := mockBookRepo.Called()
	return args.Error(0)
}
func (mockBookRepo *mockBookRepository) ReturnBook(id int) error {
	args := mockBookRepo.Called()
	return args.Error(0)
}
func NewBookRepositoryMock() *mockBookRepository {
	return &mockBookRepository{}
}
