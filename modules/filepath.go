package modules

import (
	"errors"
	"path/filepath"
)

func FilePath(fileName string) (string, error) {

	ext := filepath.Ext(fileName)

	if ext == "" {
		return "", errors.New("error")
	}

	return ext, nil
}
