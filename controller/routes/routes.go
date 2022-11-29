package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/hugovallada/rest-to-graph/controller"
	_ "github.com/hugovallada/rest-to-graph/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Post("/query", controller.GetDataFromGraphAPI)
	router.Get("/ready", controller.Ready)
	router.Get("/health", controller.Health)
	router.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8081/swagger/doc.json")))
	return router
}
