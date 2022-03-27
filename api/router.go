package api

import (
	FormAPI "upload-form/form"

	"github.com/gorilla/mux"
)

// LoadRouter - get mux router with all the routes
func LoadRouter() *mux.Router {
	router := mux.NewRouter()

	FormAPI.LoadFormRoutes(router)

	return router
}
