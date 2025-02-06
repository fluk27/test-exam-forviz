package routers

import (
	"test-exam-forviz/internal/handlers"
	"test-exam-forviz/internal/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRouter(bookSvc services.BookService) *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())
	//book
	bookHandle := handlers.NewBookHandlers(bookSvc)
	api := e.Group("/book")
	api.POST("/create", bookHandle.CreateBookHandler)
	api.GET("/list", bookHandle.SearchBooksHandler)
	api.GET("/summary", bookHandle.GetMostBorrowedBooksHandler)
	api.GET("/:id", bookHandle.GetBookByIDHandler)
	api.PUT("/:id", bookHandle.UpdateBookHandler)
	api.DELETE("/:id", bookHandle.DeleteBookHandler)
	api.PATCH("/borrow/:id", bookHandle.BorrowBookHandler)
	api.PATCH("/return/:id", bookHandle.ReturnBookHandler)
	return e
}
