package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/hugovallada/rest-to-graph/client/graphql"
	"github.com/hugovallada/rest-to-graph/files"
)

func GetDataFromGraphAPI(w http.ResponseWriter, r *http.Request) {
	multiPartData, err := parseMultiFormData(r)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid File!", 400)
		return
	}
	queryData, err := files.ReadFile(multiPartData.File)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid File!", 400)
		return
	}
	data, err := graphql.ExecuteGraphQLCall(multiPartData.URL, multiPartData.Headers, queryData)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal error while processing...", 500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	status := "ok"
	json.NewEncoder(w).Encode(status)
}

func Ready(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	status := "ok"
	json.NewEncoder(w).Encode(status)
}

func Version(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	status := "1.0"
	json.NewEncoder(w).Encode(status)
}

func parseMultiFormData(r *http.Request) (MultiPartFormData, error) {
	r.ParseMultipartForm(30 << 20)
	url := r.Form.Get("url")
	headers := r.Form.Get("headers")
	file, _, err := r.FormFile("file")
	return New(url, headers, file), err
}
