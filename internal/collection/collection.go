package collection

import (
	"fmt"
	"os"
	"path"

	"github.com/lflahive/romy-db/internal/config"
)

type CollectionInformation struct {
	Size int64
	Name string
}

func Create(name string) error {
	if collectionExists(name) {
		return fmt.Errorf("collection '%s' already exists", name)
	}

	collectionPath := getFilePath(name)

	if err := os.Mkdir(collectionPath, os.ModePerm); err != nil {
		return err
	}

	return nil
}

func Get(name string) (CollectionInformation, error) {
	if !collectionExists(name) {
		return CollectionInformation{}, fmt.Errorf("collection '%s' does not exist", name)
	}

	collectionPath := getFilePath(name)

	info, err := os.Stat(collectionPath)
	if err != nil {
		return CollectionInformation{}, err
	}

	return CollectionInformation{Size: info.Size(), Name: info.Name()}, nil
}

func collectionExists(name string) bool {
	collectionPath := getFilePath(name)

	if _, err := os.Stat(collectionPath); os.IsNotExist(err) {
		return false
	}

	return true
}

func getFilePath(name string) string {
	return path.Join(config.Configuration.StoragePath, name)
}
