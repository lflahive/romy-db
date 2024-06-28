package partition

import (
	"os"
	"testing"

	"github.com/lflahive/romy-db/internal/collection"
)

var testCollection = "test"
var testPartition = "test"

func cleanup() {
	os.RemoveAll(testCollection)
}

func TestCreatePartition(t *testing.T) {
	collection.Create(testCollection)
	err := Create(testCollection, testPartition)

	if err != nil {
		t.Error(err)
	}

	cleanup()
}

func TestCreateDuplicatePartition(t *testing.T) {
	collection.Create(testCollection)
	Create(testCollection, testPartition)
	err := Create(testCollection, testPartition)

	if err == nil {
		t.Error("Duplicate partition should return error")
	}

	cleanup()
}

func TestCreatePartitionInNonExistentCollection(t *testing.T) {
	err := Create(testCollection, testPartition)

	if err == nil {
		t.Error("Creating a partition in a non-existent collection should return error")
	}

	cleanup()
}

func TestGetPartition(t *testing.T) {
	collection.Create(testCollection)
	Create(testCollection, testCollection)
	_, err := Get(testCollection, testCollection)

	if err != nil {
		t.Error("Getting existent partition should not return error")
	}

	cleanup()
}

func TestGetNonExistentCollection(t *testing.T) {
	collection.Create(testCollection)
	_, err := Get(testCollection, testCollection)

	if err == nil {
		t.Error("Getting partition in non-existent collection should return error")
	}

	cleanup()
}

func TestGetNonExistentPartition(t *testing.T) {
	collection.Create(testCollection)
	_, err := Get(testCollection, testCollection)

	if err == nil {
		t.Error("Getting non-existent partition should return error")
	}

	cleanup()
}
