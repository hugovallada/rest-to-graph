package main

import (
	"net/http"

	"github.com/hugovallada/rest-to-graph/controller"
)

func main() {
	http.HandleFunc("/", controller.GetDataFromGraphAPI)
	http.ListenAndServe(":8080", nil)
}
