package main

import (
	"encoding/json"
	"os"
	"testing"
)

func TestBookStruct(t *testing.T) {
	b := Book{
		Id:       "1",
		Title:    "Atomic Habits",
		Author:   "James Clear",
		Price:    "300",
		Imageurl: "http://example.com/atomic.png",
	}

	data, err := json.Marshal(b)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}

	var got Book
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}

	if got != b {
		t.Errorf("got %+v, want %+v", got, b)
	}
}

// make sure the image_url json tag is wired up correctly
func TestImageURLTag(t *testing.T) {
	raw := `{"id":"1","title":"Test","author":"Someone","price":"10","image_url":"http://img.example.com/1.png"}`

	var b Book
	if err := json.Unmarshal([]byte(raw), &b); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}

	if b.Imageurl != "http://img.example.com/1.png" {
		t.Errorf("image_url not mapped correctly, got: %q", b.Imageurl)
	}
}

func TestSaveBooks(t *testing.T) {
	dir := t.TempDir()
	orig, _ := os.Getwd()
	defer os.Chdir(orig)
	os.Chdir(dir)

	books := []Book{
		{"1", "Atomic Habits", "James Clear", "300", "http://example.com/ah.png"},
	}

	if err := saveBooks(books); err != nil {
		t.Fatalf("saveBooks returned error: %v", err)
	}

	data, err := os.ReadFile("books.json")
	if err != nil {
		t.Fatalf("books.json not created: %v", err)
	}

	var saved []Book
	if err := json.Unmarshal(data, &saved); err != nil {
		t.Fatalf("saved file is not valid JSON: %v", err)
	}

	if len(saved) != 1 || saved[0].Id != "1" {
		t.Errorf("unexpected saved data: %+v", saved)
	}
}

func TestSaveBooksOverwrite(t *testing.T) {
	dir := t.TempDir()
	orig, _ := os.Getwd()
	defer os.Chdir(orig)
	os.Chdir(dir)

	first := []Book{{"1", "Book One", "Author A", "100", "http://example.com/1.png"}}
	saveBooks(first)

	second := []Book{{"2", "Book Two", "Author B", "200", "http://example.com/2.png"}}
	saveBooks(second)

	data, _ := os.ReadFile("books.json")
	var got []Book
	json.Unmarshal(data, &got)

	if len(got) != 1 || got[0].Id != "2" {
		t.Errorf("expected only second save to remain, got: %+v", got)
	}
}
