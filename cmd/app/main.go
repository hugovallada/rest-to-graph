package main

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", CallGraphQL)
	http.ListenAndServe(":8080", nil)
}

func splitHeaders(headers string) map[string]string {
	properties := make(map[string]string)

	props := strings.Split(headers, ";")

	for _, prop := range props {
		keyValue := strings.Split(prop, "=")
		properties[keyValue[0]] = keyValue[1]
	}

	return properties
}

func mountGraphQLQuery(queryBody string) map[string]string {
	return map[string]string{"query": queryBody}
}

func CallGraphQL(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(30 << 20)
	url := r.Form.Get("url")
	headers := r.Form.Get("headers")
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Invalid file", 400)
		return
	}
	queryData, err := ReadFile(file)
	if err != nil {
		http.Error(w, "Invalid file", 400)
		return
	}
	jsonData, err := json.Marshal(mountGraphQLQuery(queryData))
	if err != nil {
		http.Error(w, "Internal error while parsing the query", 500)
		return
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		http.Error(w, "Internal error while preparing the request", 500)
		return
	}
	if headers != "" {
		properties := splitHeaders(headers)
		for k, v := range properties {
			req.Header.Set(k, v)
		}
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, "Internal error while calling the api", 500)
		return
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Internal error while parsing the returned data", 500)
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
