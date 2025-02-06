package handlers

import (
	"net/http"
	"test-exam-forviz/errs"

	"github.com/labstack/echo/v4"
)

type BookHandler interface {
	CreateBookHandler(c echo.Context) error
	UpdateBookHandler(c echo.Context) error
	DeleteBookHandler(c echo.Context) error
	GetBookByIDHandler(c echo.Context) error
	SearchBooksHandler(c echo.Context) error
	GetMostBorrowedBooksHandler(c echo.Context) error
	BorrowBookHandler(c echo.Context) error
	ReturnBookHandler(c echo.Context) error
}

func HandlerError(err error) *echo.HTTPError {
	switch e := err.(type) {
	case errs.AppError:
		return echo.NewHTTPError(e.Code, e.Message)
	default:
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
}
