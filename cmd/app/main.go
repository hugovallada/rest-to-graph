package main

import (
	"log"
	"net/http"
	"time"

	"github.com/hugovallada/rest-to-graph/controller"
)

const (
	PORT = "8080"
)

func main() {
	time.Sleep(30 * time.Second)
	http.HandleFunc("/", controller.GetDataFromGraphAPI)
	http.HandleFunc("/ready", controller.Ready)
	http.HandleFunc("/health", controller.Health)
	http.HandleFunc("/version", controller.Version)
	log.Printf("Iniciando o servidor na porta %s", PORT)
	http.ListenAndServe(":"+PORT, nil)
}
