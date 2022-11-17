package main

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
)

func main() {
	http.HandleFunc("/", CallGraphQL)
	http.ListenAndServe(":8081", nil)
}

func CallGraphQL(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(30 << 20)
	url := r.Form.Get("url")
	file, _, err := r.FormFile("file")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorMessage := "Invalid file"
		json.NewEncoder(w).Encode(errorMessage)
	}

	queryData, err := ReadFile(file)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorMessage := "Invalid file"
		json.NewEncoder(w).Encode(errorMessage)
	}

	jsonQuery := map[string]string{
		"query": queryData,
	}

	jsonData, err := json.Marshal(jsonQuery)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorMessage := "Internal error while parsing the query"
		json.NewEncoder(w).Encode(errorMessage)
	}

	req, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorMessage := "Internal error while calling the graphql api"
		json.NewEncoder(w).Encode(errorMessage)
	}

	defer req.Body.Close()

	data, err := io.ReadAll(req.Body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorMessage := "Internal error while parsing the returned data"
		json.NewEncoder(w).Encode(errorMessage)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)

}

func ReadFile(file multipart.File) (string, error) {
	defer file.Close()

	data, err := io.ReadAll(file)

	if err != nil {
		return "", err
	}

	return string(data), nil
}
