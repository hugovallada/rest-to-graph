package controller

import "mime/multipart"

type MultiPartFormData struct {
	URL     string
	Headers string
	File    multipart.File
}

func New(url, headers string, file multipart.File) MultiPartFormData {
	return MultiPartFormData{URL: url, Headers: headers, File: file}
}
