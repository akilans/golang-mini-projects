package models

import (
	"errors"
	"log"

	"github.com/akilans/fiber-book-rest/initializers"
	"gorm.io/gorm"
)

// Book Type -> Books table
type Book struct {
	ID       int     `json:"id"`
	Title    string  `json:"title" validate:"required,min=1,max=100" gorm:"unique"`
	Author   string  `json:"author" validate:"required,min=1,max=50"`
	Price    float64 `json:"price" validate:"required,number"`
	ImageURL string  `json:"image_url" validate:"required,min=1,max=350"`
}

var db *gorm.DB

// Init function to connect to DB, get DB connection and migrate tables
func init() {
	initializers.LoadEnvs()
	initializers.ConnectDB()
	db = initializers.GetDB()
	SyncDB()
}

// add a book
// Get books
func GetBooks() ([]Book, error) {
	var books []Book
	result := db.Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}

// add a book
func AddBook(book Book) (id int, err error) {
	result := db.Create(&book)
	if result.Error != nil {
		return 0, result.Error
	} else {
		return book.ID, nil
	}
}

// Update a book
func UpdateBook(book Book) (id int, err error) {
	result := db.Save(&book)
	if result.Error != nil {
		return 0, result.Error
	} else {
		return book.ID, nil
	}
}

// Get book details
func GetBookByID(id int) Book {
	book, isExists := IsBookExists(id)
	if isExists {
		return book
	}
	return Book{}
}

// check book exits or not
func IsBookExists(id int) (Book, bool) {
	var book = Book{ID: id}
	result := db.Limit(1).Find(&book)
	if result.RowsAffected > 0 {
		return book, true
	}
	return Book{}, false
}

// Delete a book by id
func DeleteBookByID(id int) error {
	book, isExists := IsBookExists(id)
	if isExists {
		db.Delete(&book)
		return nil
	}
	return errors.New("Book doesn't exists")

}

// Migrate tables
func SyncDB() {
	log.Println("Start of DB migration")
	err := db.AutoMigrate(&Book{}, &User{})
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("End of DB migration")
}
