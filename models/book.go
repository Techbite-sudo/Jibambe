package models

import (
	"Jibambe/graph/model"
	// "gorm.io/gorm"
)

type Book struct {
	Base
	Title         string `json:"title"`
	Author        string `json:"author"`
	PublishedYear int    `json:"publishedYear"`
}

func (b *Book) ToGraphData() *model.Book {
	return &model.Book{
		ID:            b.ID.String(),
		Title:         b.Title,
		Author:        b.Author,
		PublishedYear: b.PublishedYear,
	}
}

