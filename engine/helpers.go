package engine

import (
	"Jibambe/models"
	"Jibambe/utils"
	"errors"
	"gorm.io/gorm/clause"
)

func FetchBookById(id string) (*models.Book, error) {
    var book models.Book
    err := utils.DB.Preload(clause.Associations).Where("id = ?", id).First(&book).Error
    if err != nil {
        return nil, errors.New("book not found")
    }
    return &book, nil
}

func FetchAllBooks() ([]models.Book, error) {
    var books []models.Book
    err := utils.DB.Preload(clause.Associations).Find(&books).Error
    if err != nil {
        return nil, errors.New("books could not be found")
    }
    return books, nil
}




