package dev

import (
	"path"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

// LoadEnvFromFile loads environment variables from a file relative to the caller's path.
func LoadEnvFromFile(file string) error {
	currdir, err := getCurrentDir(1)
	if err != nil {
		return errors.Wrap(err, "cannot load environment")
	}

	return godotenv.Load(path.Clean(path.Join(currdir, file)))
}
