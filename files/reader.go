package files

import (
	"io"
	"mime/multipart"
)

func ReadFile(file multipart.File) (string, error) {
	defer file.Close()
	data, err := io.ReadAll(file)
	return string(data), err
}
