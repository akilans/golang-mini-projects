package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// struct based on books.json file. Please refer
type Book struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Price    string `json:"price"`
	Imageurl string `json:"image_url"`
}

// define port
const PORT string = ":8080"

// message to send as json response
type Message struct {
	Msg string
}

// response as json format
func jsonMessageByte(msg string) []byte {
	errrMessage := Message{msg}
	byteContent, _ := json.Marshal(errrMessage)
	return byteContent
}

// print logs in console
func checkError(err error) {
	if err != nil {
		log.Printf("Error - %v", err)
	}

}

// main function starts here
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
	// stop the app is any error to start the server
	if err != nil {
		log.Fatal(err)
	}
}

// List all the books handler
func handleGetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := getBooks()

	// send server error as response
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
	// get book id from URL
	bookId := query.Get("id")
	book, _, err := getBookById(bookId)
	// send server error as response
	if err != nil {
		log.Printf("Server Error %v\n", err)
		w.WriteHeader(500)
		w.Write(jsonMessageByte("Internal server error"))
	} else {
		// check requested book exists or not
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
	// check for post method
	if r.Method != "POST" {
		w.WriteHeader(405)
		w.Write(jsonMessageByte(r.Method + " - Method not allowed"))
	} else {
		// read the body
		newBookByte, err := ioutil.ReadAll(r.Body)
		// check for valid data from client
		if err != nil {
			log.Printf("Client Error %v\n", err)
			w.WriteHeader(400)
			w.Write(jsonMessageByte("Bad Request"))
		} else {
			books, _ := getBooks() // get all books
			var newBooks []Book    // to add new book

			json.Unmarshal(newBookByte, &newBooks) // new book added
			books = append(books, newBooks...)     // add both
			// Write all the books in books.json file
			err = saveBooks(books)
			// send server error as response
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
	// check for post method
	if r.Method != "POST" {
		w.WriteHeader(405)
		w.Write(jsonMessageByte(r.Method + " - Method not allowed"))
	} else {
		// read the body
		updateBookByte, err := ioutil.ReadAll(r.Body)
		// check for valid data from client
		if err != nil {
			log.Printf("Client Error %v\n", err)
			w.WriteHeader(400)
			w.Write(jsonMessageByte("Bad Request"))
		} else {
			var updateBook Book // to update a book

			err = json.Unmarshal(updateBookByte, &updateBook) // new book added
			checkError(err)
			id := updateBook.Id

			book, _, _ := getBookById(id)
			// check requested book exists or not
			if (Book{}) == book {
				w.Write(jsonMessageByte("Book Not found"))
			} else {
				books, _ := getBooks()

				for i, book := range books {
					if book.Id == updateBook.Id {
						books[i] = updateBook
					}
				}
				// write books in books.json
				err = saveBooks(books)
				// send server error as response
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
	// get book id from URL
	bookId := query.Get("id")
	book, book_index, err := getBookById(bookId)
	// send server error as response
	if err != nil {
		log.Printf("Server Error %v\n", err)
		w.WriteHeader(500)
		w.Write(jsonMessageByte("Internal server error"))
	} else {
		// check requested book exists or not
		if (Book{}) == book {
			w.Write(jsonMessageByte("Book Not found"))
		} else {
			books, _ := getBooks()
			// remove books from slice
			books = append(books[:book_index], books[book_index+1:]...)
			saveBooks(books)
			w.Write(jsonMessageByte("Book deleted successfully"))
		}
	}
}
