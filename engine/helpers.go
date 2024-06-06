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



// var (
// 	ErrWrongCreds = errors.New("invalid email or password")
// 	ErrWrongToken = errors.New("invalid token")
// )

// func FetchUserByEmail(email string) (*models.User, error) {
// 	var user models.User
// 	err := utils.DB.Where("email =?", email).First(&user).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &user, nil
// }
// func FetchUserById(id string) (*models.User, error) {
// 	var user models.User
// 	err := utils.DB.Where("id =?", id).First(&user).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &user, nil
// }
// func FetchUserByAuthToken(jwt string) (*models.User, error) {
// 	userId, err := utils.ValidateJWTForAuthId(jwt)
// 	if err != nil {
// 		return nil, ErrWrongToken
// 	}
// 	return FetchUserById(userId)
// }




