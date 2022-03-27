package form

import (
	"os"
	constval "upload-form/constant"

	"github.com/gorilla/mux"
)

// LoadFormRoutes - load all form routes with form prefix
func LoadFormRoutes(router *mux.Router) {
	formRoutes := router.PathPrefix("/form").Subrouter()

	// create image directory
	os.Mkdir(constval.ImageFolderPath, os.ModePerm)

	formRoutes.HandleFunc("", FormGet).Methods("GET")
	formRoutes.HandleFunc("/upload", FormUpload).Methods("POST")

}
