package main

import (
	"log"
	"net/http"

	"github.com/hugovallada/rest-to-graph/controller/routes"
)

const (
	PORT = "8081"
)

func main() {
	log.Printf("Iniciando o servidor na porta %s\n", PORT)
	http.ListenAndServe(":"+PORT, routes.Routes())
}
