package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

// Test helper to create a temporary books.json file
func setupTestFile(t *testing.T, books []Book) (string, func()) {
	t.Helper()

	// Create a temp directory for test files
	tmpDir, err := os.MkdirTemp("", "bookstore-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}

	// Save original working directory
	originalWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get working directory: %v", err)
	}

	// Change to temp directory
	if err := os.Chdir(tmpDir); err != nil {
		t.Fatalf("Failed to change to temp dir: %v", err)
	}

	// Create books.json in temp directory
	booksJSON, err := json.Marshal(books)
	if err != nil {
		t.Fatalf("Failed to marshal books: %v", err)
	}

	if err := os.WriteFile(filepath.Join(tmpDir, "books.json"), booksJSON, 0644); err != nil {
		t.Fatalf("Failed to write books.json: %v", err)
	}

	// Return cleanup function
	cleanup := func() {
		os.Chdir(originalWd)
		os.RemoveAll(tmpDir)
	}

	return tmpDir, cleanup
}

// Sample test data
func getSampleBooks() []Book {
	return []Book{
		{Id: "1", Title: "Test Book 1", Author: "Author 1", Price: "100", Imageurl: "http://example.com/1.jpg"},
		{Id: "2", Title: "Test Book 2", Author: "Author 2", Price: "200", Imageurl: "http://example.com/2.jpg"},
		{Id: "3", Title: "Test Book 3", Author: "Author 3", Price: "300", Imageurl: "http://example.com/3.jpg"},
	}
}

func TestGetBooks(t *testing.T) {
	sampleBooks := getSampleBooks()
	_, cleanup := setupTestFile(t, sampleBooks)
	defer cleanup()

	books := getBooks()

	if len(books) != len(sampleBooks) {
		t.Errorf("Expected %d books, got %d", len(sampleBooks), len(books))
	}

	// Verify first book
	if books[0].Id != "1" {
		t.Errorf("Expected first book Id '1', got '%s'", books[0].Id)
	}
	if books[0].Title != "Test Book 1" {
		t.Errorf("Expected first book Title 'Test Book 1', got '%s'", books[0].Title)
	}
}

func TestGetBooksEmpty(t *testing.T) {
	_, cleanup := setupTestFile(t, []Book{})
	defer cleanup()

	books := getBooks()

	if len(books) != 0 {
		t.Errorf("Expected 0 books, got %d", len(books))
	}
}

func TestSaveBooks(t *testing.T) {
	_, cleanup := setupTestFile(t, []Book{})
	defer cleanup()

	newBooks := []Book{
		{Id: "10", Title: "New Book", Author: "New Author", Price: "999", Imageurl: "http://example.com/new.jpg"},
	}

	err := saveBooks(newBooks)
	if err != nil {
		t.Errorf("saveBooks returned error: %v", err)
	}

	// Read back and verify
	books := getBooks()
	if len(books) != 1 {
		t.Errorf("Expected 1 book after save, got %d", len(books))
	}
	if books[0].Id != "10" {
		t.Errorf("Expected saved book Id '10', got '%s'", books[0].Id)
	}
}

func TestSaveBooksMultiple(t *testing.T) {
	sampleBooks := getSampleBooks()
	_, cleanup := setupTestFile(t, []Book{})
	defer cleanup()

	err := saveBooks(sampleBooks)
	if err != nil {
		t.Errorf("saveBooks returned error: %v", err)
	}

	books := getBooks()
	if len(books) != 3 {
		t.Errorf("Expected 3 books after save, got %d", len(books))
	}
}

func TestBookStruct(t *testing.T) {
	book := Book{
		Id:       "1",
		Title:    "Test Title",
		Author:   "Test Author",
		Price:    "100",
		Imageurl: "http://test.com/image.jpg",
	}

	// Test JSON marshaling
	jsonBytes, err := json.Marshal(book)
	if err != nil {
		t.Errorf("Failed to marshal book: %v", err)
	}

	// Test JSON unmarshaling
	var unmarshaledBook Book
	err = json.Unmarshal(jsonBytes, &unmarshaledBook)
	if err != nil {
		t.Errorf("Failed to unmarshal book: %v", err)
	}

	if unmarshaledBook.Id != book.Id {
		t.Errorf("Expected Id '%s', got '%s'", book.Id, unmarshaledBook.Id)
	}
	if unmarshaledBook.Title != book.Title {
		t.Errorf("Expected Title '%s', got '%s'", book.Title, unmarshaledBook.Title)
	}
}

func TestBookJSONTags(t *testing.T) {
	// Test that JSON tags work correctly (image_url vs Imageurl)
	jsonStr := `{"id":"1","title":"Test","author":"Auth","price":"50","image_url":"http://img.com"}`

	var book Book
	err := json.Unmarshal([]byte(jsonStr), &book)
	if err != nil {
		t.Errorf("Failed to unmarshal with JSON tags: %v", err)
	}

	if book.Imageurl != "http://img.com" {
		t.Errorf("Expected Imageurl 'http://img.com', got '%s'", book.Imageurl)
	}
}

func TestGetBooksReturnsCorrectOrder(t *testing.T) {
	sampleBooks := getSampleBooks()
	_, cleanup := setupTestFile(t, sampleBooks)
	defer cleanup()

	books := getBooks()

	for i, book := range books {
		if book.Id != sampleBooks[i].Id {
			t.Errorf("Book at index %d has wrong Id. Expected '%s', got '%s'", i, sampleBooks[i].Id, book.Id)
		}
	}
}

func TestSaveBooksOverwrite(t *testing.T) {
	sampleBooks := getSampleBooks()
	_, cleanup := setupTestFile(t, sampleBooks)
	defer cleanup()

	// Save new books (should overwrite)
	newBooks := []Book{
		{Id: "99", Title: "Only Book", Author: "Solo Author", Price: "999", Imageurl: "http://solo.com"},
	}

	err := saveBooks(newBooks)
	if err != nil {
		t.Errorf("saveBooks returned error: %v", err)
	}

	books := getBooks()
	if len(books) != 1 {
		t.Errorf("Expected 1 book after overwrite, got %d", len(books))
	}
	if books[0].Id != "99" {
		t.Errorf("Expected overwritten book Id '99', got '%s'", books[0].Id)
	}
}

// Benchmark tests
func BenchmarkGetBooks(b *testing.B) {
	// Create temp file with sample data
	tmpDir, err := os.MkdirTemp("", "bookstore-bench")
	if err != nil {
		b.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	originalWd, _ := os.Getwd()
	os.Chdir(tmpDir)
	defer os.Chdir(originalWd)

	sampleBooks := getSampleBooks()
	booksJSON, _ := json.Marshal(sampleBooks)
	os.WriteFile("books.json", booksJSON, 0644)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = getBooks()
	}
}

func BenchmarkSaveBooks(b *testing.B) {
	tmpDir, err := os.MkdirTemp("", "bookstore-bench")
	if err != nil {
		b.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	originalWd, _ := os.Getwd()
	os.Chdir(tmpDir)
	defer os.Chdir(originalWd)

	sampleBooks := getSampleBooks()
	booksJSON, _ := json.Marshal(sampleBooks)
	os.WriteFile("books.json", booksJSON, 0644)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = saveBooks(sampleBooks)
	}
}
