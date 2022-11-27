package routes

import (
	"net/http"

	"github.com/hugovallada/rest-to-graph/controller"
)

func Routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", controller.GetDataFromGraphAPI)
	mux.HandleFunc("/ready", controller.Ready)
	mux.HandleFunc("/health", controller.Health)
	return mux
}
