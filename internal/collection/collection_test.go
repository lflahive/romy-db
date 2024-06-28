package collection

import (
	"os"
	"path"
	"testing"
)

var testCollection = "test"

func cleanup() {
	os.Remove(testCollection)
}

func TestCreateCollection(t *testing.T) {
	err := Create(testCollection)

	if err != nil {
		t.Error(err)
	}

	cleanup()
}

func TestCreateDuplicateCollection(t *testing.T) {
	Create(testCollection)
	err := Create(testCollection)

	if err == nil {
		t.Error("Duplicate collection should return error")
	}

	cleanup()
}

func TestCreateCollectionInNonExistentRootPath(t *testing.T) {
	err := Create(path.Join("long/path", testCollection))

	if err == nil {
		t.Error("Creating a collection in a non-existent root path should return error")
	}

	cleanup()
}

func TestGetCollection(t *testing.T) {
	Create(testCollection)
	_, err := Get(testCollection)

	if err != nil {
		t.Error("Getting existent collection should not return error")
	}

	cleanup()
}

func TestGetNonExistentCollection(t *testing.T) {
	_, err := Get(testCollection)

	if err == nil {
		t.Error("Getting non-existent collection should return error")
	}

	cleanup()
}
