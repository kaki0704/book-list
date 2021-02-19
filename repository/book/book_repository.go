package bookrepository

import (
	"book-list/models"
	"log"

	"gorm.io/gorm"
)

// BookRepository is...
type BookRepository struct{}

func logfatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// GetBooks function is...
func (b BookRepository) GetBooks(db *gorm.DB) ([]models.Book, error) {
	var books []models.Book
	if err := db.Find(&books).Error; err != nil {
		return books, err
	}
	return books, nil
}

// GetBook is
func (b BookRepository) GetBook(db *gorm.DB, id int) (models.Book, error) {
	var book models.Book
	if err := db.Where("Id = ?", id).Find(&book).Error; err != nil {
		return book, err
	}
	return book, nil
}

// AddBook is
func (b BookRepository) AddBook(db *gorm.DB, book models.Book) (int, error) {
	if err := db.Create(&book).Error; err != nil {
		return 0, err
	}
	return book.ID, nil
}

// UpdateBook is
func (b BookRepository) UpdateBook(db *gorm.DB, book models.Book) (int64, error) {
	if err := db.Model(&book).Updates(models.Book{Title: book.Title, Author: book.Author, Year: book.Year}).Error; err != nil {
		return 0, err
	}
	return int64(book.ID), nil
}

// RemoveBook is
func (b BookRepository) RemoveBook(db *gorm.DB, id int) (int64, error) {
	if err := db.Delete(&models.Book{}, id).Error; err != nil {
		return 0, err
	}
	return int64(id), nil
}
