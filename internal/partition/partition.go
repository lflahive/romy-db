package partition

import (
	"fmt"
	"os"
	"path"

	"github.com/lflahive/romy-db/internal/collection"
	"github.com/lflahive/romy-db/internal/config"
)

type PartitionInformation struct {
	Size int64
	Name string
}

func Create(collectionName string, name string) error {
	_, err := collection.Get(collectionName)
	if err != nil {
		return err
	}

	if partitionExists(collectionName, name) {
		return fmt.Errorf("partition '%s' in collection '%s' already exists", name, collectionName)
	}

	partitionPath := getFilePath(collectionName, name)

	if err := os.Mkdir(partitionPath, os.ModePerm); err != nil {
		return err
	}

	return nil
}

func Get(collectionName string, name string) (PartitionInformation, error) {
	_, err := collection.Get(collectionName)
	if err != nil {
		return PartitionInformation{}, err
	}

	if !partitionExists(collectionName, name) {
		return PartitionInformation{}, fmt.Errorf("partition '%s' in collection '%s' does not exist", name, collectionName)
	}

	partitionPath := getFilePath(collectionName, name)

	info, err := os.Stat(partitionPath)
	if err != nil {
		return PartitionInformation{}, err
	}

	return PartitionInformation{Size: info.Size(), Name: info.Name()}, nil
}

func partitionExists(collectionName string, name string) bool {
	partitionPath := getFilePath(collectionName, name)

	if _, err := os.Stat(partitionPath); os.IsNotExist(err) {
		return false
	}

	return true
}

func getFilePath(collectionName string, name string) string {
	return path.Join(config.Configuration.StoragePath, collectionName, name)
}
