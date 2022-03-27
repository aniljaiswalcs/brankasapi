package form

import (
	"os"

	"github.com/gorilla/mux"
)

const (
	ImageFolderPath = "image"
)

// LoadFormRoutes - load all form routes with form prefix
func LoadFormRoutes(router *mux.Router) {
	formRoutes := router.PathPrefix("/form").Subrouter()

	// create image directory
	os.Mkdir(ImageFolderPath, os.ModePerm)

	formRoutes.HandleFunc("", FormGet).Methods("GET")
	formRoutes.HandleFunc("/upload", FormUpload).Methods("POST")

}
