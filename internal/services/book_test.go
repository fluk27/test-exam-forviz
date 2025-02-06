package services_test

import (
	"errors"
	"test-exam-forviz/config"
	"test-exam-forviz/constant"
	"test-exam-forviz/internal/models"
	"test-exam-forviz/internal/repositories/db"
	"test-exam-forviz/internal/services"
	"test-exam-forviz/loggers"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestBorrowBook(t *testing.T) {

	loggers.InitLogger(config.App{Env: "dev"})
	testCases := []struct {
		name          string
		requestId     int
		mockData      models.BookRepository
		expectSuccess models.BookResponse
		expectError   error
	}{
		{
			name:      "TestBorrowBookSuccess",
			requestId: 1,
			mockData: models.BookRepository{
				ID:         1,
				Title:      "title test2",
				Author:     "author test2",
				Category:   "category test2",
				IsBorrowed: false,
			},

			expectSuccess: models.BookResponse{
				Message: constant.BookBorrowSuccessMessage,
				Data:    nil,
			},
			expectError: nil,
		},
		{
			name:      "TestBorrowBookErrorInternalServerError",
			requestId: 1,
			mockData: models.BookRepository{
				ID:         1,
				Title:      "title test2",
				Author:     "author test2",
				Category:   "category test2",
				IsBorrowed: false,
			},

			expectSuccess: models.BookResponse{
				Message: constant.BookBorrowSuccessMessage,
				Data:    nil,
			},
			expectError: errors.New(constant.BookErrorMessageInternalServerError),
		},
		{
			name:      "TestBorrowBookFindNotFound",
			requestId: 1,
			mockData: models.BookRepository{
				ID:         1,
				Title:      "title test2",
				Author:     "author test2",
				Category:   "category test2",
				IsBorrowed: false,
			},

			expectSuccess: models.BookResponse{
				Message: constant.BookBorrowSuccessMessage,
				Data:    nil,
			},
			expectError: errors.New(constant.BookErrorsMessageFindNotFound),
		},
		{
			name:      "TestBorrowBookNotMatchFindNotFound",
			requestId: 1,
			mockData: models.BookRepository{
				ID:         1,
				Title:      "title test2",
				Author:     "author test2",
				Category:   "category test2",
				IsBorrowed: false,
			},

			expectSuccess: models.BookResponse{
				Message: constant.BookBorrowSuccessMessage,
				Data:    nil,
			},
			expectError: errors.New(constant.BookErrorMessageInternalServerError),
		},
		{
			name:      "TestBorrowBookBorrowed",
			requestId: 1,
			mockData: models.BookRepository{
				ID:         1,
				Title:      "title test2",
				Author:     "author test2",
				Category:   "category test2",
				IsBorrowed: true,
			},

			expectSuccess: models.BookResponse{
				Message: constant.BookBorrowSuccessMessage,
				Data:    nil,
			},
			expectError: errors.New(constant.BookBarrowErrorMessage),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			bookRepo := db.NewBookRepositoryMock()

			switch tC.name {
			case "TestBorrowBookFindNotFound":
				bookRepo.On("FindByID").Return(tC.mockData, gorm.ErrRecordNotFound)
				bookRepo.On("BorrowBook").Return(tC.expectError)
				break
			case "TestBorrowBookNotMatchFindNotFound":
				bookRepo.On("FindByID").Return(tC.mockData, errors.New(""))
				bookRepo.On("BorrowBook").Return(tC.expectError)
				break
			case "TestBorrowBookBorrowed":
				bookRepo.On("FindByID").Return(tC.mockData, nil)
				bookRepo.On("BorrowBook").Return(tC.expectError)
				break
			case "TestBorrowBookErrorInternalServerError":
				bookRepo.On("FindByID").Return(tC.mockData, nil)
				bookRepo.On("BorrowBook").Return(tC.expectError)
				break
			default:
				bookRepo.On("FindByID").Return(tC.mockData, nil)
				bookRepo.On("BorrowBook").Return(nil)
				break
			}

			bookSvc := services.NewBookService(bookRepo)

			resp, err := bookSvc.BorrowBook(tC.requestId)
			if err != nil {
				assert.EqualError(t, tC.expectError, err.Error())
			} else {

				assert.Equal(t, tC.expectSuccess, resp)
			}

		})
	}
}

func TestReturnBook(t *testing.T) {

	loggers.InitLogger(config.App{Env: "dev"})
	testCases := []struct {
		name          string
		requestId     int
		mockData      models.BookRepository
		expectSuccess models.BookResponse
		expectError   error
	}{
		{
			name:      "TestReturnBookSuccess",
			requestId: 1,
			mockData: models.BookRepository{
				ID:         1,
				Title:      "title test2",
				Author:     "author test2",
				Category:   "category test2",
				IsBorrowed: true,
			},

			expectSuccess: models.BookResponse{
				Message: constant.BookReturnSuccessMessage,
				Data:    nil,
			},
			expectError: nil,
		},
		{
			name:      "TestReturnBookErrorInternalServerError",
			requestId: 1,
			mockData: models.BookRepository{
				ID:         1,
				Title:      "title test2",
				Author:     "author test2",
				Category:   "category test2",
				IsBorrowed: true,
			},

			expectSuccess: models.BookResponse{
				Message: constant.BookReturnSuccessMessage,
				Data:    nil,
			},
			expectError: errors.New(constant.BookErrorMessageInternalServerError),
		},
		{
			name:      "TestReturnBookFindNotFound",
			requestId: 1,
			mockData: models.BookRepository{
				ID:         1,
				Title:      "title test2",
				Author:     "author test2",
				Category:   "category test2",
				IsBorrowed: true,
			},

			expectSuccess: models.BookResponse{
				Message: constant.BookReturnSuccessMessage,
				Data:    nil,
			},
			expectError: errors.New(constant.BookErrorsMessageFindNotFound),
		},
		{
			name:      "TestReturnBookNotMatchFindNotFound",
			requestId: 1,
			mockData: models.BookRepository{
				ID:         1,
				Title:      "title test2",
				Author:     "author test2",
				Category:   "category test2",
				IsBorrowed: true,
			},

			expectSuccess: models.BookResponse{
				Message: constant.BookReturnSuccessMessage,
				Data:    nil,
			},
			expectError: errors.New(constant.BookErrorMessageInternalServerError),
		},
		{
			name:      "TestReturnBookReturned",
			requestId: 1,
			mockData: models.BookRepository{
				ID:         1,
				Title:      "title test2",
				Author:     "author test2",
				Category:   "category test2",
				IsBorrowed: false,
			},

			expectSuccess: models.BookResponse{
				Message: constant.BookReturnSuccessMessage,
				Data:    nil,
			},
			expectError: errors.New(constant.BookReturnErrorMessage),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			bookRepo := db.NewBookRepositoryMock()

			switch tC.name {
			case "TestReturnBookFindNotFound":
				bookRepo.On("FindByID").Return(tC.mockData, gorm.ErrRecordNotFound)
				bookRepo.On("ReturnBook").Return(tC.expectError)
				break
			case "TestReturnBookNotMatchFindNotFound":
				bookRepo.On("FindByID").Return(tC.mockData, errors.New(""))
				bookRepo.On("ReturnBook").Return(tC.expectError)
				break
			case "TestReturnBookReturned":
				bookRepo.On("FindByID").Return(tC.mockData, nil)
				bookRepo.On("ReturnBook").Return(tC.expectError)
				break
			case "TestReturnBookErrorInternalServerError":
				bookRepo.On("FindByID").Return(tC.mockData, nil)
				bookRepo.On("ReturnBook").Return(tC.expectError)
				break
			default:
				bookRepo.On("FindByID").Return(tC.mockData, nil)
				bookRepo.On("ReturnBook").Return(nil)
				break
			}

			bookSvc := services.NewBookService(bookRepo)

			resp, err := bookSvc.ReturnBook(tC.requestId)
			if err != nil {
				assert.EqualError(t, tC.expectError, err.Error())
			} else {

				assert.Equal(t, tC.expectSuccess, resp)
			}

		})
	}
}

func TestCreateBook(t *testing.T) {

	loggers.InitLogger(config.App{Env: "dev"})
	testCases := []struct {
		name          string
		request       models.BookRequest
		mockData      models.BookRepository
		expectSuccess models.BookResponse
		expectError   error
	}{
		{
			name: "createBookSuccess",
			request: models.BookRequest{
				Title:    "title test2",
				Author:   "author test2",
				Category: "category test2",
			},
			mockData: models.BookRepository{
				Title:    "title test2",
				Author:   "author test2",
				Category: "category test2",
			},

			expectSuccess: models.BookResponse{
				Message: constant.BookCreateSuccessMessage,
				Data:    nil,
			},
			expectError: nil,
		},
		{
			name: "createBookError",
			request: models.BookRequest{
				Title:    "Ttitle test2",
				Author:   "author test2",
				Category: "category test2",
			},
			mockData: models.BookRepository{
				Title:    "title test2",
				Author:   "author test2",
				Category: "category test2",
			},

			expectSuccess: models.BookResponse{
				Message: constant.BookCreateSuccessMessage,
				Data:    nil,
			},
			expectError: errors.New(constant.BookErrorMessageInternalServerError),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {

			bookRepo := db.NewBookRepositoryMock()
			if tC.expectError != nil {
				bookRepo.On("Create").Return(errors.New(constant.BookErrorMessageInternalServerError))
			} else {
				bookRepo.On("Create").Return(nil)
			}

			bookSvc := services.NewBookService(bookRepo)
			resp, err := bookSvc.CreateBook(tC.request)
			if err != nil {
				assert.EqualError(t, tC.expectError, err.Error())
			} else {

				assert.Equal(t, tC.expectSuccess, resp)
			}

		})
	}
}
func TestDeleteBook(t *testing.T) {

	loggers.InitLogger(config.App{Env: "dev"})
	testCases := []struct {
		name          string
		requestId     int
		mockData      models.BookRepository
		expectSuccess models.BookResponse
		expectError   error
	}{
		{
			name:      "TestDeleteBookSuccess",
			requestId: 1,
			mockData: models.BookRepository{
				ID:         1,
				Title:      "title test2",
				Author:     "author test2",
				Category:   "category test2",
				IsBorrowed: true,
			},

			expectSuccess: models.BookResponse{
				Message: constant.BookDeleteSuccessMessage,
				Data:    nil,
			},
			expectError: nil,
		},
		{
			name:      "TestDeleteBookErrorInternalServerError",
			requestId: 1,
			mockData: models.BookRepository{
				ID:         1,
				Title:      "title test2",
				Author:     "author test2",
				Category:   "category test2",
				IsBorrowed: true,
			},

			expectSuccess: models.BookResponse{
				Message: constant.BookDeleteSuccessMessage,
				Data:    nil,
			},
			expectError: errors.New(constant.BookErrorMessageInternalServerError),
		},
		{
			name:      "TestDeleteBookFindNotFound",
			requestId: 1,
			mockData: models.BookRepository{
				ID:         1,
				Title:      "title test2",
				Author:     "author test2",
				Category:   "category test2",
				IsBorrowed: true,
			},

			expectSuccess: models.BookResponse{
				Message: constant.BookDeleteSuccessMessage,
				Data:    nil,
			},
			expectError: errors.New(constant.BookErrorsMessageFindNotFound),
		},
		{
			name:      "TestDeleteBookNotMatchFindNotFound",
			requestId: 1,
			mockData: models.BookRepository{
				ID:         1,
				Title:      "title test2",
				Author:     "author test2",
				Category:   "category test2",
				IsBorrowed: true,
			},

			expectSuccess: models.BookResponse{
				Message: constant.BookDeleteSuccessMessage,
				Data:    nil,
			},
			expectError: errors.New(constant.BookErrorMessageInternalServerError),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			bookRepo := db.NewBookRepositoryMock()

			switch tC.name {
			case "TestDeleteBookFindNotFound":
				bookRepo.On("FindByID").Return(tC.mockData, gorm.ErrRecordNotFound)
				bookRepo.On("Delete").Return(tC.expectError)
				break
			case "TestDeleteBookNotMatchFindNotFound":
				bookRepo.On("FindByID").Return(tC.mockData, errors.New(""))
				bookRepo.On("Delete").Return(tC.expectError)
				break
			case "TestDeleteBookErrorInternalServerError":
				bookRepo.On("FindByID").Return(tC.mockData, nil)
				bookRepo.On("Delete").Return(tC.expectError)
				break
			default:
				bookRepo.On("FindByID").Return(tC.mockData, nil)
				bookRepo.On("Delete").Return(nil)
				break
			}

			bookSvc := services.NewBookService(bookRepo)

			resp, err := bookSvc.DeleteBook(tC.requestId)
			if err != nil {
				assert.EqualError(t, tC.expectError, err.Error())
			} else {

				assert.Equal(t, tC.expectSuccess, resp)
			}

		})
	}
}
func TestGetBookByID(t *testing.T) {
	const dateFormat = "02/01/2006"
	loggers.InitLogger(config.App{Env: "dev"})
	testCases := []struct {
		name          string
		requestId     int
		mockData      models.BookRepository
		expectSuccess models.BookResponse
		expectError   error
	}{
		{
			name:      "TestGetBookByIDSuccess",
			requestId: 1,
			mockData: models.BookRepository{
				ID:          1,
				Title:       "title test2",
				Author:      "author test2",
				Category:    "category test2",
				IsBorrowed:  true,
				BorrowCount: 1,
				CreateAt:    time.Now(),
				UpdateAt:    time.Now(),
			},

			expectSuccess: models.BookResponse{
				Message: constant.BookGetSuccessMessage,
				Data: &models.BookData{
					ID:          1,
					Title:       "title test2",
					Author:      "author test2",
					Category:    "category test2",
					IsBorrowed:  true,
					BorrowCount: 1,
					CreateAt:    time.Now().Format(dateFormat),
					UpdateAt:    time.Now().Format(dateFormat),
				},
			},
			expectError: nil,
		},
		{
			name:      "TestGetBookByIDErrorInternalServerError",
			requestId: 1,
			mockData: models.BookRepository{
				ID:          1,
				Title:       "title test2",
				Author:      "author test2",
				Category:    "category test2",
				IsBorrowed:  true,
				BorrowCount: 1,
				CreateAt:    time.Now(),
				UpdateAt:    time.Now(),
			},

			expectSuccess: models.BookResponse{
				Message: constant.BookGetSuccessMessage,
				Data:    nil,
			},
			expectError: errors.New(constant.BookErrorMessageInternalServerError),
		},
		{
			name:      "TestGetBookByIDFindNotFound",
			requestId: 1,
			mockData: models.BookRepository{
				ID:          0,
				Title:       "title test2",
				Author:      "author test2",
				Category:    "category test2",
				IsBorrowed:  true,
				BorrowCount: 1,
				CreateAt:    time.Now(),
				UpdateAt:    time.Now(),
			},

			expectSuccess: models.BookResponse{
				Message: constant.BookGetSuccessMessage,
				Data:    nil,
			},
			expectError: errors.New(constant.BookErrorsMessageFindNotFound),
		},
		{
			name:      "TestGetBookByIDNotMatchFindNotFound",
			requestId: 1,
			mockData: models.BookRepository{
				ID:          1,
				Title:       "title test2",
				Author:      "author test2",
				Category:    "category test2",
				IsBorrowed:  true,
				BorrowCount: 1,
				CreateAt:    time.Now(),
				UpdateAt:    time.Now(),
			},

			expectSuccess: models.BookResponse{
				Message: constant.BookGetSuccessMessage,
				Data:    nil,
			},
			expectError: errors.New(constant.BookErrorMessageInternalServerError),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			bookRepo := db.NewBookRepositoryMock()

			switch tC.name {
			case "TestGetBookByIDFindNotFound":
				bookRepo.On("FindByID").Return(tC.mockData, gorm.ErrRecordNotFound)
				break
			case "TestGetBookByIDNotMatchFindNotFound":
				bookRepo.On("FindByID").Return(tC.mockData, errors.New(""))

				break
			case "TestGetBookByIDErrorInternalServerError":
				bookRepo.On("FindByID").Return(tC.mockData, errors.New(""))

				break
			default:
				bookRepo.On("FindByID").Return(tC.mockData, nil)
				break
			}

			bookSvc := services.NewBookService(bookRepo)

			resp, err := bookSvc.GetBookByID(tC.requestId)
			if err != nil {
				assert.EqualError(t, tC.expectError, err.Error())
			} else {

				assert.Equal(t, tC.expectSuccess, resp)
			}

		})
	}
}
func TestGetMostBorrowedBooks(t *testing.T) {
	const dateFormat = "02/01/2006"
	loggers.InitLogger(config.App{Env: "dev"})
	testCases := []struct {
		name          string
		mockData      []models.BookRepository
		expectSuccess models.BookListResponse
		expectError   error
	}{
		{
			name: "TestGetMostBorrowedBooksByIDSuccess",
			mockData: []models.BookRepository{
				{
					ID:          1,
					Title:       "title test2",
					Author:      "author test2",
					Category:    "category test2",
					IsBorrowed:  true,
					BorrowCount: 10,
					CreateAt:    time.Now(),
					UpdateAt:    time.Now(),
				},
				{
					ID:          2,
					Title:       "title test2",
					Author:      "author test2",
					Category:    "category test2",
					IsBorrowed:  true,
					BorrowCount: 1,
					CreateAt:    time.Now(),
					UpdateAt:    time.Now(),
				},
			},

			expectSuccess: models.BookListResponse{
				Message: constant.BookGetSuccessMessage,
				Data: []models.BookData{
					{
						ID:          1,
						Title:       "title test2",
						Author:      "author test2",
						Category:    "category test2",
						IsBorrowed:  true,
						BorrowCount: 10,
						CreateAt:    time.Now().Format(dateFormat),
						UpdateAt:    time.Now().Format(dateFormat),
					},
					{
						ID:          2,
						Title:       "title test2",
						Author:      "author test2",
						Category:    "category test2",
						IsBorrowed:  true,
						BorrowCount: 1,
						CreateAt:    time.Now().Format(dateFormat),
						UpdateAt:    time.Now().Format(dateFormat),
					},
				},
			},
			expectError: nil,
		},
		{
			name: "TestGetMostBorrowedBooksByIDErrorInternalServerError",
			mockData: []models.BookRepository{
				{
					ID:          1,
					Title:       "title test2",
					Author:      "author test2",
					Category:    "category test2",
					IsBorrowed:  true,
					BorrowCount: 10,
					CreateAt:    time.Now(),
					UpdateAt:    time.Now(),
				},
				{
					ID:          2,
					Title:       "title test2",
					Author:      "author test2",
					Category:    "category test2",
					IsBorrowed:  true,
					BorrowCount: 1,
					CreateAt:    time.Now(),
					UpdateAt:    time.Now(),
				},
			},

			expectSuccess: models.BookListResponse{
				Message: constant.BookGetSuccessMessage,
				Data:    nil,
			},
			expectError: errors.New(constant.BookErrorMessageInternalServerError),
		},
		{
			name: "TestGetMostBorrowedBooksByIDFindNotFound",
			mockData: []models.BookRepository{
				{
					ID:          1,
					Title:       "title test2",
					Author:      "author test2",
					Category:    "category test2",
					IsBorrowed:  true,
					BorrowCount: 10,
					CreateAt:    time.Now(),
					UpdateAt:    time.Now(),
				},
				{
					ID:          2,
					Title:       "title test2",
					Author:      "author test2",
					Category:    "category test2",
					IsBorrowed:  true,
					BorrowCount: 1,
					CreateAt:    time.Now(),
					UpdateAt:    time.Now(),
				},
			},

			expectSuccess: models.BookListResponse{
				Message: constant.BookGetSuccessMessage,
				Data:    nil,
			},
			expectError: errors.New(constant.BookErrorsMessageFindNotFound),
		},
		{
			name: "TestGetMostBorrowedBooksByIDNotMatchFindNotFound",
			mockData: []models.BookRepository{
				{
					ID:          1,
					Title:       "title test2",
					Author:      "author test2",
					Category:    "category test2",
					IsBorrowed:  true,
					BorrowCount: 10,
					CreateAt:    time.Now(),
					UpdateAt:    time.Now(),
				},
				{
					ID:          2,
					Title:       "title test2",
					Author:      "author test2",
					Category:    "category test2",
					IsBorrowed:  true,
					BorrowCount: 1,
					CreateAt:    time.Now(),
					UpdateAt:    time.Now(),
				},
			},

			expectSuccess: models.BookListResponse{
				Message: constant.BookGetSuccessMessage,
				Data:    nil,
			},
			expectError: errors.New(constant.BookErrorMessageInternalServerError),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			bookRepo := db.NewBookRepositoryMock()

			switch tC.name {
			case "TestGetMostBorrowedBooksByIDFindNotFound":
				bookRepo.On("FindAll").Return(tC.mockData, gorm.ErrRecordNotFound)
				break
			case "TestGetMostBorrowedBooksByIDNotMatchFindNotFound":
				bookRepo.On("FindAll").Return(tC.mockData, errors.New(""))

				break
			case "TestGetMostBorrowedBooksByIDErrorInternalServerError":
				bookRepo.On("FindAll").Return(tC.mockData, errors.New(""))

				break
			default:
				bookRepo.On("FindAll").Return(tC.mockData, nil)
				break
			}

			bookSvc := services.NewBookService(bookRepo)

			resp, err := bookSvc.GetMostBorrowedBooks()
			if err != nil {
				assert.EqualError(t, tC.expectError, err.Error())
			} else {

				assert.Equal(t, tC.expectSuccess, resp)
			}

		})
	}
}
func TestSearchBooks(t *testing.T) {
	const dateFormat = "02/01/2006"
	loggers.InitLogger(config.App{Env: "dev"})
	testCases := []struct {
		name          string
		title         string
		author        string
		category      string
		mockData      []models.BookRepository
		expectSuccess models.BookListResponse
		expectError   error
	}{
		{
			name:     "TestSearchBooksByIDSuccess",
			title:    "",
			author:   "",
			category: "",
			mockData: []models.BookRepository{
				{
					ID:          1,
					Title:       "title test2",
					Author:      "author test2",
					Category:    "category test2",
					IsBorrowed:  true,
					BorrowCount: 10,
					CreateAt:    time.Now(),
					UpdateAt:    time.Now(),
				},
				{
					ID:          2,
					Title:       "title test2",
					Author:      "author test2",
					Category:    "category test2",
					IsBorrowed:  true,
					BorrowCount: 1,
					CreateAt:    time.Now(),
					UpdateAt:    time.Now(),
				},
			},

			expectSuccess: models.BookListResponse{
				Message: constant.BookGetSuccessMessage,
				Data: []models.BookData{
					{
						ID:          1,
						Title:       "title test2",
						Author:      "author test2",
						Category:    "category test2",
						IsBorrowed:  true,
						BorrowCount: 10,
						CreateAt:    time.Now().Format(dateFormat),
						UpdateAt:    time.Now().Format(dateFormat),
					},
					{
						ID:          2,
						Title:       "title test2",
						Author:      "author test2",
						Category:    "category test2",
						IsBorrowed:  true,
						BorrowCount: 1,
						CreateAt:    time.Now().Format(dateFormat),
						UpdateAt:    time.Now().Format(dateFormat),
					},
				},
			},
			expectError: nil,
		},
		{
			name:     "TestSearchBooksByIDErrorInternalServerError",
			title:    "",
			author:   "",
			category: "",
			mockData: []models.BookRepository{
				{
					ID:          1,
					Title:       "title test2",
					Author:      "author test2",
					Category:    "category test2",
					IsBorrowed:  true,
					BorrowCount: 10,
					CreateAt:    time.Now(),
					UpdateAt:    time.Now(),
				},
				{
					ID:          2,
					Title:       "title test2",
					Author:      "author test2",
					Category:    "category test2",
					IsBorrowed:  true,
					BorrowCount: 1,
					CreateAt:    time.Now(),
					UpdateAt:    time.Now(),
				},
			},

			expectSuccess: models.BookListResponse{
				Message: constant.BookGetSuccessMessage,
				Data:    nil,
			},
			expectError: errors.New(constant.BookErrorMessageInternalServerError),
		},
		{
			name:     "TestSearchBooksByIDFindNotFound",
			title:    "",
			author:   "",
			category: "",
			mockData: []models.BookRepository{
				{
					ID:          1,
					Title:       "title test2",
					Author:      "author test2",
					Category:    "category test2",
					IsBorrowed:  true,
					BorrowCount: 10,
					CreateAt:    time.Now(),
					UpdateAt:    time.Now(),
				},
				{
					ID:          2,
					Title:       "title test2",
					Author:      "author test2",
					Category:    "category test2",
					IsBorrowed:  true,
					BorrowCount: 1,
					CreateAt:    time.Now(),
					UpdateAt:    time.Now(),
				},
			},

			expectSuccess: models.BookListResponse{
				Message: constant.BookGetSuccessMessage,
				Data:    nil,
			},
			expectError: errors.New(constant.BookErrorsMessageFindNotFound),
		},
		{
			name:     "TestSearchBooksByIDNotMatchFindNotFound",
			title:    "",
			author:   "",
			category: "",
			mockData: []models.BookRepository{
				{
					ID:          1,
					Title:       "title test2",
					Author:      "author test2",
					Category:    "category test2",
					IsBorrowed:  true,
					BorrowCount: 10,
					CreateAt:    time.Now(),
					UpdateAt:    time.Now(),
				},
				{
					ID:          2,
					Title:       "title test2",
					Author:      "author test2",
					Category:    "category test2",
					IsBorrowed:  true,
					BorrowCount: 1,
					CreateAt:    time.Now(),
					UpdateAt:    time.Now(),
				},
			},

			expectSuccess: models.BookListResponse{
				Message: constant.BookGetSuccessMessage,
				Data:    nil,
			},
			expectError: errors.New(constant.BookErrorMessageInternalServerError),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			bookRepo := db.NewBookRepositoryMock()

			switch tC.name {
			case "TestSearchBooksByIDFindNotFound":
				bookRepo.On("FindAll").Return(tC.mockData, gorm.ErrRecordNotFound)
				break
			case "TestSearchBooksByIDNotMatchFindNotFound":
				bookRepo.On("FindAll").Return(tC.mockData, errors.New(""))

				break
			case "TestSearchBooksByIDErrorInternalServerError":
				bookRepo.On("FindAll").Return(tC.mockData, errors.New(""))

				break
			default:
				bookRepo.On("FindAll").Return(tC.mockData, nil)
				break
			}

			bookSvc := services.NewBookService(bookRepo)

			resp, err := bookSvc.SearchBooks(tC.title, tC.author, tC.category)
			if err != nil {
				assert.EqualError(t, tC.expectError, err.Error())
			} else {

				assert.Equal(t, tC.expectSuccess, resp)
			}

		})
	}
}
func TestUpdateBook(t *testing.T) {

	loggers.InitLogger(config.App{Env: "dev"})
	testCases := []struct {
		name          string
		requestId     int
		requestBody   models.BookRequest
		mockData      models.BookRepository
		expectSuccess models.BookResponse
		expectError   error
	}{
		{
			name:      "TestUpdateBookSuccess",
			requestId: 1,
			requestBody: models.BookRequest{

				Title:    "title test2",
				Author:   "author test2",
				Category: "category test2",
			},
			mockData: models.BookRepository{
				ID:         1,
				Title:      "title test2",
				Author:     "author test2",
				Category:   "category test2",
				IsBorrowed: true,
			},

			expectSuccess: models.BookResponse{
				Message: constant.BookUpdateSuccessMessage,
				Data:    nil,
			},
			expectError: nil,
		},
		{
			name:      "TestUpdateBookErrorInternalServerError",
			requestId: 1,
			requestBody: models.BookRequest{

				Title:    "title test2",
				Author:   "author test2",
				Category: "category test2",
			},
			mockData: models.BookRepository{
				ID:         1,
				Title:      "title test2",
				Author:     "author test2",
				Category:   "category test2",
				IsBorrowed: true,
			},

			expectSuccess: models.BookResponse{
				Message: constant.BookUpdateSuccessMessage,
				Data:    nil,
			},
			expectError: errors.New(constant.BookErrorMessageInternalServerError),
		},
		{
			name:      "TestUpdateBookFindNotFound",
			requestId: 2,
			requestBody: models.BookRequest{

				Title:    "title test2",
				Author:   "author test2",
				Category: "category test2",
			},
			mockData: models.BookRepository{},

			expectSuccess: models.BookResponse{
				Message: constant.BookUpdateSuccessMessage,
				Data:    nil,
			},
			expectError: errors.New(constant.BookErrorsMessageFindNotFound),
		},
		{
			name:      "TestUpdateBookNotMatchFindNotFound",
			requestId: 1,
			requestBody: models.BookRequest{

				Title:    "title test2",
				Author:   "author test2",
				Category: "category test2",
			},
			mockData: models.BookRepository{},

			expectSuccess: models.BookResponse{
				Message: constant.BookUpdateSuccessMessage,
				Data:    nil,
			},
			expectError: errors.New(constant.BookErrorMessageInternalServerError),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			bookRepo := db.NewBookRepositoryMock()

			switch tC.name {
			case "TestUpdateBookFindNotFound":
				bookRepo.On("FindByID").Return(tC.mockData, gorm.ErrRecordNotFound)
				bookRepo.On("Update").Return(tC.expectError)
				break
			case "TestUpdateBookNotMatchFindNotFound":
				bookRepo.On("FindByID").Return(tC.mockData, errors.New(""))
				bookRepo.On("Update").Return(tC.expectError)
				break
			case "TestUpdateBookErrorInternalServerError":
				bookRepo.On("FindByID").Return(tC.mockData, nil)
				bookRepo.On("Update").Return(tC.expectError)
				break
			default:
				bookRepo.On("FindByID").Return(tC.mockData, nil)
				bookRepo.On("Update").Return(nil)
				break
			}

			bookSvc := services.NewBookService(bookRepo)

			resp, err := bookSvc.UpdateBook(tC.requestId, tC.requestBody)
			if err != nil {
				assert.EqualError(t, tC.expectError, err.Error())
			} else {

				assert.Equal(t, tC.expectSuccess, resp)
			}

		})
	}
}
