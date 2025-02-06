package db

import (
	"fmt"
	"test-exam-forviz/internal/models"

	"gorm.io/gorm"
)

type bookRepository struct {
	db *gorm.DB
}

// BorrowBook implements BookRepository.
func (b bookRepository) BorrowBook(id, count int) error {
	err := b.db.Transaction(func(tx *gorm.DB) error {
		db := tx.Model(&models.BookRepository{}).Where("id=?", id).Update("is_borrowed", true).Update("borrow_count", count)
		if db.Error != nil {
			return db.Error
		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

// Create implements BookRepository.
func (b bookRepository) Create(book models.BookRepository) error {
	err := b.db.Transaction(func(tx *gorm.DB) error {
		db := tx.Create(&book)
		if db.Error != nil {
			return db.Error
		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

// Delete implements BookRepository.
func (b bookRepository) Delete(id int) error {
	err := b.db.Transaction(func(tx *gorm.DB) error {
		db := tx.Where("id = ?", id).Delete(&models.BookRepository{})
		if db.Error != nil {
			return db.Error
		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

// FindAll implements BookRepository.
func (b bookRepository) FindAll(title, author, category, sortName, sortType string) ([]models.BookRepository, error) {
	bookList := []models.BookRepository{}
	query := b.db
	if title != "" {
		query = query.Where("title LIKE ?", "%"+title+"%")

	}
	if author != "" {
		query = query.Where("author LIKE %?%", "%"+author+"%")

	}
	if category != "" {
		query = query.Where("category LIKE %?%", "%"+category+"%")

	}
	if sortName != "" && sortType != "" {

		switch sortType {
		case "asc":
			query = query.Order(fmt.Sprintf("%s asc", sortName))
			break
		case "desc":
			query = query.Order(fmt.Sprintf("%s desc", sortName))
			break
		}
	}
	db := query.Find(&bookList)
	if db.Error != nil {
		return bookList, db.Error
	}
	return bookList, nil
}

// FindByID implements BookRepository.
func (b bookRepository) FindByID(id int) (models.BookRepository, error) {
	bookRepoResp := models.BookRepository{}
	db := b.db.Where("id = ?", id).First(&bookRepoResp)
	if db.Error != nil {
		return bookRepoResp, db.Error
	}
	return bookRepoResp, nil
}

// ReturnBook implements BookRepository.
func (b bookRepository) ReturnBook(id int) error {
	err := b.db.Transaction(func(tx *gorm.DB) error {
		db := tx.Model(&models.BookRepository{}).Where("id=?", id).Update("is_borrowed", false)
		if db.Error != nil {
			return db.Error
		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

// Update implements BookRepository.
func (b bookRepository) Update(req models.BookRepository) error {
	err := b.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id", req.ID).Updates(&req).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return bookRepository{db: db}
}
