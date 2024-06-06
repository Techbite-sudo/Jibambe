package books

import (
	"Jibambe/engine"
	"Jibambe/graph/model"
	"Jibambe/models"
	"Jibambe/utils"
	"errors"
)

func CreateBook(input model.BookInput) (*model.Book, error) {
	newBook := models.Book{
		Title:         input.Title,
		Author:        input.Author,
		PublishedYear: input.PublishedYear,
	}
	err := utils.DB.Create(&newBook).Error
	if err != nil {
		return nil, errors.New("book could not be created")
	}
	return newBook.ToGraphData(), nil
}

func UpdateBook(id string, input model.BookInput) (*model.Book, error) {
	book, err := engine.FetchBookById(id)
	if err != nil {
		return nil, errors.New("book not found")
	}
	book.Title = input.Title
	book.Author = input.Author
	book.PublishedYear = input.PublishedYear
	err = utils.DB.Save(&book).Error
	if err != nil {
		return nil, errors.New("book could not be saved")
	}
	return book.ToGraphData(), nil
}

func DeleteBook(id string) (bool, error) {
	book, err := engine.FetchBookById(id)
	if err != nil {
		return false, errors.New("book not found")
	}
	err = utils.DB.Delete(&book).Error
	if err != nil {
		return false, errors.New("book could not be deleted")
	}
	return true, nil
}

func Books() ([]*model.Book, error) {
	books, err := engine.FetchAllBooks()
	if err != nil {
		return nil, errors.New("books could not be fetched")
	}
	var graphBooks []*model.Book
	for _, book := range books {
		graphBooks = append(graphBooks, book.ToGraphData())
	}
	return graphBooks, nil
}

func Book(id string) (*model.Book, error) {
	book, err := engine.FetchBookById(id)
	if err != nil {
		return nil, errors.New("book not found")
	}
	return book.ToGraphData(), nil
}
