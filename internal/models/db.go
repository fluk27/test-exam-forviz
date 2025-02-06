package models

import "time"

type BookRepository struct {
	ID          int       `gorm:"primaryKey;autoIncrement"`
	Title       string    `gorm:"index;not null"`
	Author      string    `gorm:"index;not null"`
	Category    string    `gorm:"index;not null"`
	IsBorrowed  bool      `gorm:"default:false"`
	BorrowCount int       `gorm:"borrow_count;default:0"`
	UpdateAt    time.Time `gorm:"autoCreateTime"`
	CreateAt    time.Time `gorm:"autoUpdateTime"`
}
