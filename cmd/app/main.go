package main

import (
	"log"
	"net/http"

	"github.com/hugovallada/rest-to-graph/controller"
)

const (
	PORT = "8080"
)

func main() {
	http.HandleFunc("/", controller.GetDataFromGraphAPI)
	http.HandleFunc("/ready", controller.Ready)
	http.HandleFunc("/health", controller.Health)
	log.Printf("Iniciando o servidor na porta %s", PORT)
	http.ListenAndServe(":"+PORT, nil)
}
