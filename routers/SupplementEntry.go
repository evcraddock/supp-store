package routers

import (
	"github.com/evcraddock/supp-store/controllers"
	"github.com/gorilla/mux"
)

func SetSupplementEntryRoutes(router *mux.Router) *mux.Router {

	router.HandleFunc("/api/entries", controllers.CreateEntry).Methods("POST")
	router.HandleFunc("/api/entries", controllers.GetEntries).Methods("GET")

	return router
}
