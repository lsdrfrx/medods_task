package storage

import (
	"testing"
)

func TestCreate(t *testing.T) {
	db := NewDB()
	err := db.Open()
	if err != nil {
		t.Fatal("Unable to open storage")
	}

	err = db.Create("John")

	if err != nil {
		t.Fatalf("Unable to create new document: %s", err.Error())
	}
}

func TestDelete(t *testing.T) {
	db := NewDB()
	err := db.Open()
	if err != nil {
		t.Fatal("Unable to open storage")
	}

	err = db.Delete("John")
	if err != nil {
		t.Fatalf("Unable to delete document: %s", err.Error())
	}

	res, err := db.Get("John")
	if err == nil || res.UserId != "" {
		t.Fatal("Error or result is nil")
	}
}

func TestUpdate(t *testing.T) {
	db := NewDB()
	err := db.Open()
	if err != nil {
		t.Fatal("Unable to open storage")
	}

	err = db.Create("Peter")
	if err != nil {
		t.Fatalf("Unable to create document: %s", err.Error())
	}

	err = db.Update("Peter", "12345")
	if err != nil {
		t.Fatalf("Unable to update document: %s", err.Error())
	}

	res, err := db.Get("Peter")
	if res.RefreshToken != "12345" {
		t.Fatal("Update didn't performed")
	}
	if err != nil {
		t.Fatalf("Unable to get document: %s", err.Error())
	}
}
