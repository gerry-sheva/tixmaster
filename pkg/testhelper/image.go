package testhelper

import (
	"os"
	"path/filepath"
)

func NewImage() (*os.File, error) {
	img, err := os.Open(filepath.Join("..", "testhelper", "data", "ado.jpg"))

	if err != nil {
		return nil, err
	}

	return img, nil
}
