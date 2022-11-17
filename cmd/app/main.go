package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", CallGraphQL)
	http.ListenAndServe(":8081", nil)
}

func CallGraphQL(w http.ResponseWriter, r *http.Request) {
	queryData, err := ReadFile("develop/mock/query.graphql")

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

	req, err := http.Post("https://api.spacex.land/graphql/", "application/json", bytes.NewBuffer(jsonData))

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

func ReadFile(fileName string) (string, error) {
	file, err := os.Open(fileName)

	if err != nil {
		return "", err
	}

	defer file.Close()

	data, err := io.ReadAll(file)

	if err != nil {
		return "", err
	}

	return string(data), nil
}
