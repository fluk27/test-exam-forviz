package handlers

import (
	"net/http"
	"regexp"
	"strconv"
	"test-exam-forviz/internal/models"
	"test-exam-forviz/internal/services"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type bookHandlers struct {
	service services.BookService
}

// BorrowBookHandler implements BookHandler.
func (b bookHandlers) BorrowBookHandler(c echo.Context) error {
	paramsId := c.Param("id")
	resultId := regexp.MustCompile(`^[1-9][0-9]*$`).MatchString(paramsId)
	if !resultId {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": "id must have digit only and start 1"})
	}

	id, err := strconv.Atoi(paramsId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	bookResp, err := b.service.BorrowBook(id)
	if err != nil {
		return HandlerError(err)
	}
	return c.JSONPretty(http.StatusOK, bookResp, "")
}

// CreateBookHandler implements BookHandler.
func (b bookHandlers) CreateBookHandler(c echo.Context) error {
	bookReq := new(models.BookRequest)
	if err := c.Bind(bookReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	validate := validator.New()
	if err := validate.Struct(bookReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	bookResp, err := b.service.CreateBook(*bookReq)
	if err != nil {
		return HandlerError(err)
	}
	return c.JSONPretty(http.StatusCreated, bookResp, "")
}

// DeleteBookHandler implements BookHandler.
func (b bookHandlers) DeleteBookHandler(c echo.Context) error {
	paramsId := c.Param("id")
	resultId := regexp.MustCompile(`^[1-9][0-9]*$`).MatchString(paramsId)
	if !resultId {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": "id must have digit only and start 1"})
	}

	id, err := strconv.Atoi(paramsId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	bookResp, err := b.service.DeleteBook(id)
	if err != nil {
		return HandlerError(err)
	}
	return c.JSONPretty(http.StatusOK, bookResp, "")
}

// GetBookByIDHandler implements BookHandler.
func (b bookHandlers) GetBookByIDHandler(c echo.Context) error {

	Id := c.Param("id")
	resultId := regexp.MustCompile(`^[1-9][0-9]*$`).MatchString(Id)
	if !resultId {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": "id must have digit only and start 1"})
	}
	id, err := strconv.Atoi(Id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	bookResp, err := b.service.GetBookByID(id)
	if err != nil {
		return HandlerError(err)
	}
	return c.JSONPretty(http.StatusOK, bookResp, "")
}

// GetMostBorrowedBooksHandler implements BookHandler.
func (b bookHandlers) GetMostBorrowedBooksHandler(c echo.Context) error {
	bookResp, err := b.service.GetMostBorrowedBooks()
	if err != nil {
		return HandlerError(err)
	}
	return c.JSONPretty(http.StatusOK, bookResp, "")
}

// ReturnBookHandler implements BookHandler.
func (b bookHandlers) ReturnBookHandler(c echo.Context) error {
	paramsId := c.Param("id")
	resultId := regexp.MustCompile(`^[1-9][0-9]*$`).MatchString(paramsId)
	if !resultId {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": "id must have digit only and start 1"})
	}

	id, err := strconv.Atoi(paramsId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	bookResp, err := b.service.ReturnBook(id)
	if err != nil {
		return HandlerError(err)
	}
	return c.JSONPretty(http.StatusOK, bookResp, "")
}

// SearchBooksHandler implements BookHandler.
func (b bookHandlers) SearchBooksHandler(c echo.Context) error {
	title := c.QueryParam("title")
	author := c.QueryParam("author")
	category := c.QueryParam("category")
	bookResp, err := b.service.SearchBooks(title, author, category)
	if err != nil {
		return HandlerError(err)
	}
	return c.JSONPretty(http.StatusOK, bookResp, "")
}

// UpdateBookHandler implements BookHandler.
func (b bookHandlers) UpdateBookHandler(c echo.Context) error {
	paramsId := c.Param("id")
	resultId := regexp.MustCompile(`^[1-9][0-9]*$`).MatchString(paramsId)
	if !resultId {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": "id must have digit only and start 1"})
	}
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	bookReq := new(models.BookRequest)
	if err = c.Bind(bookReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	validate := validator.New()
	if err := validate.Struct(bookReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	bookResp, err := b.service.UpdateBook(id, *bookReq)
	if err != nil {
		return HandlerError(err)
	}
	return c.JSONPretty(http.StatusOK, bookResp, "")
}

func NewBookHandlers(service services.BookService) BookHandler {
	return bookHandlers{service: service}
}
