package main

import (
	"encoding/json"
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
