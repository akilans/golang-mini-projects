package main

import (
	"flag"
	"fmt"
	"os"
)

// struct based on books.json file. Please refer
type Book struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Price    string `json:"price"`
	Imageurl string `json:"image_url"`
}

func main() {
	/*
		get books --all or --id
		./bookstore get --all
		./bookstore get --id 5
	*/
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)
	getAll := getCmd.Bool("all", false, "List all the books")
	getId := getCmd.String("id", "", "Get book by id")

	/*
		add a book with id ,title, author, price, image_url
		./bookstore add --id 6 --title test-book --author akilan --price 200 --image_url http://akilan.com/test.png
	*/

	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addId := addCmd.String("id", "", "Book id")
	addTitle := addCmd.String("title", "", "Book title")
	addAuthor := addCmd.String("author", "", "Book author")
	addPrice := addCmd.String("price", "", "Book price")
	addImageUrl := addCmd.String("image_url", "", "Book image URL")

	/*
		update a book with id ,title, author, price, image_url
		./bookstore update --id 6 --title test-book-1 --author akilan1 --price 2001 --image_url http://akilan.com/test.png1
	*/

	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
	updateId := updateCmd.String("id", "", "Book id")
	updateTitle := updateCmd.String("title", "", "Book title")
	updateAuthor := updateCmd.String("author", "", "Book author")
	updatePrice := updateCmd.String("price", "", "Book price")
	updateImageUrl := updateCmd.String("image_url", "", "Book image URL")

	/*
		delete a book by --id
		./bookstore delete --id 6
	*/
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	deleteId := deleteCmd.String("id", "", "Delete book by id")

	// validation
	if len(os.Args) < 2 {
		fmt.Println("Expected get, add, update, delete commands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "get":
		handleGetBooks(getCmd, getAll, getId)
	case "add":
		handleAddBook(addCmd, addId, addTitle, addAuthor, addPrice, addImageUrl, true)
	case "update":
		handleAddBook(updateCmd, updateId, updateTitle, updateAuthor, updatePrice, updateImageUrl, false)
	case "delete":
		handleDeleteBook(deleteCmd, deleteId)
	default:
		fmt.Println("Please provide get, update, update, delete commands")
		os.Exit(1)
	}

}
