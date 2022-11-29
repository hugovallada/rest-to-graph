package main

import (
	"log"
	"net/http"

	"github.com/hugovallada/rest-to-graph/controller/routes"
)

const (
	PORT = "8081"
)

// @title  Rest To Graph
// @version 1.0
// @description Application that allows other apps to call a graphql via a rest endpoint
// @termsOfService http://swagger.io/terms

// @contact.name Hugo Lopes
// @contact.url https://github.com/hugovallada
// @contact.email valladahugo@gmail.com

// @license.name Apache 2.0

// @host localhost:8081
// @BasePath /
func main() {
	log.Printf("Iniciando o servidor na porta %s\n", PORT)
	http.ListenAndServe(":"+PORT, routes.Routes())
}
