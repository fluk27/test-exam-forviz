package models

type BookResponse struct {
	Message string    `json:"message"`
	Data    *BookData `json:"data,omitempty"`
}
type BookListResponse struct {
	Message string     `json:"message"`
	Data    []BookData `json:"data"`
}
type BookData struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Category    string `json:"category"`
	IsBorrowed  bool   `json:"is_borrowed"`
	BorrowCount int    `json:"borrow_count"`
	UpdateAt    string `json:"update_at"`
	CreateAt    string `json:"create_at"`
}
type BookRequest struct {
	Title    string `json:"title" validate:"required"`
	Author   string `json:"author" validate:"required"`
	Category string `json:"category" validate:"required"`
}
