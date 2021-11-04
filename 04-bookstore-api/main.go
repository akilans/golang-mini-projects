package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const PORT string = ":8080"

type Message struct {
	Msg string
}

func main() {

	// http://localhost:8080
	http.HandleFunc("/", handleGetBooks)

	// http://localhost:8080/book?id=1
	http.HandleFunc("/book", handleGetBookById)

	// http://localhost:8080/add
	http.HandleFunc("/add", handleAddBook)

	// http://localhost:8080/update
	http.HandleFunc("/update", handleUpdateBook)

	// http://localhost:8080/delete?id=1
	http.HandleFunc("/delete", handleDeleteBookById)

	fmt.Printf("App is listening on %v\n", PORT)
	err := http.ListenAndServe(PORT, nil)

	if err != nil {
		log.Fatal(err)
	}
}

// List all the books handler
func handleGetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := getBooks()
	if err != nil {
		log.Printf("Server Error %v\n", err)
		w.WriteHeader(500)
		w.Write(jsonMessageByte("Internal server error"))
	} else {
		booksByte, _ := json.Marshal(books)
		w.Write(booksByte)
	}

}

// get book by id handler
func handleGetBookById(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()

	bookId := query.Get("id")
	book, _, err := getBookById(bookId)

	if err != nil {
		log.Printf("Server Error %v\n", err)
		w.WriteHeader(500)
		w.Write(jsonMessageByte("Internal server error"))
	} else {
		if (Book{}) == book {
			w.Write(jsonMessageByte("Book Not found"))
		} else {
			bookByte, _ := json.Marshal(book)
			w.Write(bookByte)
		}
	}
}

// add book handler
func handleAddBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(405)
		w.Write(jsonMessageByte(r.Method + " - Method not allowed"))
	} else {
		// read the body
		newBookByte, err := ioutil.ReadAll(r.Body)

		if err != nil {
			log.Printf("Client Error %v\n", err)
			w.WriteHeader(400)
			w.Write(jsonMessageByte("Bad Request"))
		} else {
			books, _ := getBooks() // get all books
			var newBooks []Book    // to add new book

			json.Unmarshal(newBookByte, &newBooks) // new book added
			books = append(books, newBooks...)     // add both
			// Write all the books in new file
			err = saveBooks(books)

			if err != nil {
				log.Printf("Server Error %v\n", err)
				w.WriteHeader(500)
				w.Write(jsonMessageByte("Internal server error"))
			} else {
				w.Write(jsonMessageByte("New book added successfully"))
			}

		}
	}
}

// update book handler
func handleUpdateBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(405)
		w.Write(jsonMessageByte(r.Method + " - Method not allowed"))
	} else {
		// read the body
		updateBookByte, err := ioutil.ReadAll(r.Body)

		if err != nil {
			log.Printf("Client Error %v\n", err)
			w.WriteHeader(400)
			w.Write(jsonMessageByte("Bad Request"))
		} else {
			var updateBook Book // to add new book

			err = json.Unmarshal(updateBookByte, &updateBook) // new book added
			checkError(err)
			id := updateBook.Id

			book, _, _ := getBookById(id)

			if (Book{}) == book {
				w.Write(jsonMessageByte("Book Not found"))
			} else {
				books, _ := getBooks()

				for i, book := range books {
					if book.Id == updateBook.Id {
						books[i] = updateBook
					}
				}
				// write books
				err = saveBooks(books)

				if err != nil {
					log.Printf("Server Error %v\n", err)
					w.WriteHeader(500)
					w.Write(jsonMessageByte("Internal server error"))
				} else {
					w.Write(jsonMessageByte("Book updated successfully"))
				}
			}
		}
	}
}

// delete book by id handler
func handleDeleteBookById(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()

	bookId := query.Get("id")
	book, book_index, err := getBookById(bookId)

	if err != nil {
		log.Printf("Server Error %v\n", err)
		w.WriteHeader(500)
		w.Write(jsonMessageByte("Internal server error"))
	} else {

		if (Book{}) == book {
			w.Write(jsonMessageByte("Book Not found"))
		} else {
			books, _ := getBooks()
			books = append(books[:book_index], books[book_index+1:]...)
			saveBooks(books)
			w.Write(jsonMessageByte("Book deleted successfully"))
		}
	}
}
