package services

import (
	"errors"
	"test-exam-forviz/constant"
	"test-exam-forviz/errs"
	"test-exam-forviz/internal/models"
	"test-exam-forviz/internal/repositories/db"
	"test-exam-forviz/loggers"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

const dateFormat = "02/01/2006"

type bookService struct {
	repo db.BookRepository
}

// BorrowBook implements BookService.
func (b bookService) BorrowBook(id int) (models.BookResponse, error) {
	book, err := b.repo.FindByID(id)
	if err != nil {
		loggers.Error("Error FindByID book",
			zap.String("type", "repo"),
			zap.Error(err),
			zap.Int("book_id", id))
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return models.BookResponse{}, errs.NewInternalServerError(constant.BookErrorsMessageFindNotFound)
		} else {
			return models.BookResponse{}, errs.NewInternalServerError(constant.BookErrorMessageInternalServerError)
		}

	}
	if book.IsBorrowed {
		return models.BookResponse{}, errs.NewBadRequest(constant.BookBarrowErrorMessage)
	}
	err = b.repo.BorrowBook(id, book.BorrowCount+1)
	if err != nil {
		loggers.Error("Error Borrow book",
			zap.String("type", "repo"),
			zap.Error(err),
			zap.Int("book_id", id))
		return models.BookResponse{}, errs.NewInternalServerError(constant.BookErrorMessageInternalServerError)
	}
	return models.BookResponse{
		Message: constant.BookBorrowSuccessMessage,
	}, nil
}

// ReturnBook implements BookService.
func (b bookService) ReturnBook(id int) (models.BookResponse, error) {
	book, err := b.repo.FindByID(id)
	if err != nil {
		loggers.Error("Error FindByID book",
			zap.String("type", "repo"),
			zap.Error(err),
			zap.Int("book_id", id))
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return models.BookResponse{}, errs.NewInternalServerError(constant.BookErrorsMessageFindNotFound)
		} else {
			return models.BookResponse{}, errs.NewInternalServerError(constant.BookErrorMessageInternalServerError)
		}

	}
	if !book.IsBorrowed {
		return models.BookResponse{}, errs.NewBadRequest(constant.BookReturnErrorMessage)
	}
	err = b.repo.ReturnBook(id)
	if err != nil {
		loggers.Error("Error Return book",
			zap.String("type", "repo"),
			zap.Error(err),
			zap.Int("book_id", id))
		return models.BookResponse{}, errs.NewInternalServerError(constant.BookErrorMessageInternalServerError)
	}
	return models.BookResponse{
		Message: constant.BookReturnSuccessMessage,
	}, nil
}

// CreateBook implements BookService.
func (b bookService) CreateBook(book models.BookRequest) (models.BookResponse, error) {
	bookDataCreate := models.BookRepository{
		Title:    book.Title,
		Author:   book.Author,
		Category: book.Category,
	}
	err := b.repo.Create(bookDataCreate)
	if err != nil {
		loggers.Error("Error Create book",
			zap.String("type", "repo"),
			zap.Error(err),
			zap.Any("request", bookDataCreate))
		return models.BookResponse{}, errs.NewInternalServerError(constant.BookErrorMessageInternalServerError)
	}
	return models.BookResponse{
		Message: constant.BookCreateSuccessMessage,
	}, nil
}

// DeleteBook implements BookService.
func (b bookService) DeleteBook(id int) (models.BookResponse, error) {
	book, err := b.repo.FindByID(id)
	if err != nil {
		loggers.Error("Error FindByID book",
			zap.String("type", "repo"),
			zap.Error(err),
			zap.Int("book_id", id))
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return models.BookResponse{}, errs.NewInternalServerError(constant.BookErrorsMessageFindNotFound)
		} else {
			return models.BookResponse{}, errs.NewInternalServerError(constant.BookErrorMessageInternalServerError)
		}
	}
	err = b.repo.Delete(book.ID)
	if err != nil {
		loggers.Error("Error Delete book",
			zap.String("type", "repo"),
			zap.Error(err),
			zap.Int("book_id", id))
	}
	return models.BookResponse{
		Message: constant.BookDeleteSuccessMessage,
	}, nil
}

// GetBookByID implements BookService.
func (b bookService) GetBookByID(id int) (models.BookResponse, error) {
	book, err := b.repo.FindByID(id)
	if err != nil {
		loggers.Error("Error FindByID book",
			zap.String("type", "repo"),
			zap.Error(err),
			zap.Int("book_id", id))
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return models.BookResponse{}, errs.NewInternalServerError(constant.BookErrorsMessageFindNotFound)
		} else {
			return models.BookResponse{}, errs.NewInternalServerError(constant.BookErrorMessageInternalServerError)
		}
	}
	data := models.BookData{
		ID:          book.ID,
		Title:       book.Title,
		Author:      book.Author,
		Category:    book.Category,
		IsBorrowed:  book.IsBorrowed,
		BorrowCount: book.BorrowCount,
		CreateAt:    book.CreateAt.Format(dateFormat),
		UpdateAt:    book.UpdateAt.Format(dateFormat),
	}
	return models.BookResponse{
		Message: constant.BookGetSuccessMessage,
		Data:    &data,
	}, nil
}

// GetMostBorrowedBooks implements BookService.
func (b bookService) GetMostBorrowedBooks() (models.BookListResponse, error) {
	books, err := b.repo.FindAll("", "", "", "borrow_count", "desc")
	if err != nil {
		loggers.Error("Error FindAll book",
			zap.String("type", "repo"),
			zap.Error(err))
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return models.BookListResponse{}, errs.NewInternalServerError(constant.BookErrorsMessageFindNotFound)
		} else {
			return models.BookListResponse{}, errs.NewInternalServerError(constant.BookErrorMessageInternalServerError)
		}
	}
	bookList := []models.BookData{}
	for _, book := range books {
		bookData := models.BookData{
			ID:          book.ID,
			Title:       book.Title,
			Author:      book.Author,
			Category:    book.Category,
			IsBorrowed:  book.IsBorrowed,
			BorrowCount: book.BorrowCount,
			CreateAt:    book.CreateAt.Format(dateFormat),
			UpdateAt:    book.UpdateAt.Format(dateFormat),
		}
		bookList = append(bookList, bookData)
	}
	return models.BookListResponse{
		Message: constant.BookGetSuccessMessage,
		Data:    bookList,
	}, nil
}

// SearchBooks implements BookService.
func (b bookService) SearchBooks(title string, author string, category string) (models.BookListResponse, error) {
	books, err := b.repo.FindAll(title, author, category, "", "")
	if err != nil {
		loggers.Error("Error FindAll book",
			zap.String("type", "repo"),
			zap.Error(err))
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return models.BookListResponse{}, errs.NewInternalServerError(constant.BookErrorsMessageFindNotFound)
		} else {
			return models.BookListResponse{}, errs.NewInternalServerError(constant.BookErrorMessageInternalServerError)
		}
	}
	bookList := []models.BookData{}
	for _, book := range books {
		bookData := models.BookData{
			ID:          book.ID,
			Title:       book.Title,
			Author:      book.Author,
			Category:    book.Category,
			IsBorrowed:  book.IsBorrowed,
			BorrowCount: book.BorrowCount,
			CreateAt:    book.CreateAt.Format(dateFormat),
			UpdateAt:    book.UpdateAt.Format(dateFormat),
		}
		bookList = append(bookList, bookData)
	}
	return models.BookListResponse{
		Message: constant.BookGetSuccessMessage,
		Data:    bookList,
	}, nil
}

// UpdateBook implements BookService.
func (b bookService) UpdateBook(id int, book models.BookRequest) (models.BookResponse, error) {
	bookRepo, err := b.repo.FindByID(id)
	if err != nil {
		loggers.Error("Error FindAll book",
			zap.String("type", "repo"),
			zap.Error(err))
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return models.BookResponse{}, errs.NewInternalServerError(constant.BookErrorsMessageFindNotFound)
		} else {
			return models.BookResponse{}, errs.NewInternalServerError(constant.BookErrorMessageInternalServerError)
		}
	}
	bookDataUpdate := models.BookRepository{
		ID:          id,
		Title:       book.Title,
		Author:      book.Author,
		Category:    book.Category,
		IsBorrowed:  bookRepo.IsBorrowed,
		BorrowCount: bookRepo.BorrowCount,
	}
	err = b.repo.Update(bookDataUpdate)
	if err != nil {
		loggers.Error("Error Update book",
			zap.String("type", "repo"),
			zap.Error(err),
			zap.Any("request", bookDataUpdate))
		return models.BookResponse{}, errs.NewInternalServerError(constant.BookErrorMessageInternalServerError)
	}
	return models.BookResponse{
		Message: constant.BookUpdateSuccessMessage,
	}, nil
}

func NewBookService(repo db.BookRepository) BookService {
	return bookService{repo: repo}
}
