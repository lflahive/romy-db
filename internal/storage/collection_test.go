package storage

import (
	"os"
	"path"
	"testing"
)

var testCollection = "test"

func cleanup() {
	os.Remove("test")
}

func TestCreateCollection(t *testing.T) {
	err := CreateCollection(testCollection)

	if err != nil {
		t.Error(err)
	}

	cleanup()
}

func TestCreateDuplicateCollection(t *testing.T) {
	CreateCollection(testCollection)
	err := CreateCollection(testCollection)

	if err == nil {
		t.Error("Duplicate collection should return error")
	}

	cleanup()
}

func TestCreateCollectionInNonExistentRootPath(t *testing.T) {
	err := CreateCollection(path.Join("long/path", testCollection))

	if err == nil {
		t.Error("Creating a collection in a non-existent root path should return error")
	}

	cleanup()
}

func TestGetCollection(t *testing.T) {
	CreateCollection(testCollection)
	_, err := GetCollection(testCollection)

	if err != nil {
		t.Error("Getting non-existent collection should return error")
	}

	cleanup()
}

func TestGetNonExistentCollection(t *testing.T) {
	_, err := GetCollection(testCollection)

	if err == nil {
		t.Error("Getting non-existent collection should return error")
	}

	cleanup()
}
