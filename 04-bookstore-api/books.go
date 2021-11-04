package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// struct based on books.json file. Please refer
type Book struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Price    string `json:"price"`
	Imageurl string `json:"image_url"`
}

func jsonMessageByte(msg string) []byte {
	errrMessage := Message{msg}
	byteContent, _ := json.Marshal(errrMessage)
	return byteContent
}

func checkError(err error) {
	if err != nil {
		log.Printf("Error - %v", err)
	}

}

// Get books - returns books and error
func getBooks() ([]Book, error) {
	books := []Book{}
	booksByte, err := ioutil.ReadFile("./books.json")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(booksByte, &books)
	if err != nil {
		return nil, err
	}
	return books, nil
}

// Get books - returns book and error
func getBookById(id string) (Book, int, error) {
	books, err := getBooks()
	var requestedBook Book
	var requestedBookIndex int

	if err != nil {
		return Book{}, 0, err
	}

	for i, book := range books {
		if book.Id == id {
			requestedBook = book
			requestedBookIndex = i
		}
	}

	return requestedBook, requestedBookIndex, nil
}

// save books to books.json file
func saveBooks(books []Book) error {

	// converting into bytes for writing into a file
	booksBytes, err := json.Marshal(books)

	checkError(err)

	err = ioutil.WriteFile("./books.json", booksBytes, 0644)

	return err

}
