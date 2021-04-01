package tool

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

func GetFolderFrom(url string) (string, error) {
	path := filepath.Join(os.Getenv("HOME"), SaveFolder)

	absolutePath, err := filepath.Abs(filepath.Join(os.Getenv("HOME"), SaveFolder, filepath.Base(url)))
	if err != nil {
		return "", errors.WithStack(err)
	}

	// To prevent path traversal attack
	relative, err := filepath.Rel(path, absolutePath)
	if err != nil {
		return "", errors.WithStack(err)
	}

	if strings.Contains(relative, "..") {
		return "", errors.WithStack(errors.New("Your download file may have a path traversal attack"))
	}

	return absolutePath, nil
}
