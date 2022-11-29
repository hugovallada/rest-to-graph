package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/hugovallada/rest-to-graph/client/graphql"
	"github.com/hugovallada/rest-to-graph/controller/response"
	"github.com/hugovallada/rest-to-graph/files"
)

// Make a GraphQL call godoc
// @Sumary 			Calls a GraphQL endpoint
// @Description 	Receives a graphql file and calls a graphql endpoint
// @Tags 			query
// @Accept			multipart/form-data
// @Produce 		json
// @Param			url	    formData string true  "Url do endpoint graph"
// @Param 			file    formData file   true  "Arquivo .graphql"
// @Param           headers formData string false "Headers a serem enviados para o endpoint"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /query [post]
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

// Health Endpoint godoc
// @Sumary 			Health Endpoint
// @Description 	Health Endpoint for livenessProbe
// @Tags 			health
// @Produce 		json
// @Success 200 {object} string
// @Failure 500 {object} string
// @Router /health [get]
func Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	status := response.New()
	json.NewEncoder(w).Encode(status)
}

// Ready Endpoint godoc
// @Sumary 			Ready Endpoint
// @Description 	Ready Endpoint for livenessProbe
// @Tags 			health
// @Produce 		json
// @Success 200 {object} string
// @Failure 500 {object} string
// @Router /ready [get]
func Ready(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	status := response.New()
	json.NewEncoder(w).Encode(status)
}

func parseMultiFormData(r *http.Request) (MultiPartFormData, error) {
	r.ParseMultipartForm(30 << 20)
	url := r.Form.Get("url")
	headers := r.Form.Get("headers")
	file, _, err := r.FormFile("file")
	return New(url, headers, file), err
}
