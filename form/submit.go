package form

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"strings"
	CONFIG "upload-form/config"
	constval "upload-form/constant"

	utils "upload-form/util"
)

// FormGet  Get form HTML
func FormGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	var response = make(map[string]interface{})

	// parse html template
	tpl, err := template.ParseFiles(constval.UploadHTMLFile)
	if err != nil {
		fmt.Println("FormGet", err)
		utils.SetReponse(w, constval.StatusCodeServerError, err.Error(), "", response)
		return
	}

	// replace variables in template
	err = tpl.Execute(w, map[string]string{
		"Auth": CONFIG.AuthFormKey,
	})
	if err != nil {
		fmt.Println("FormGet", err)
		utils.SetReponse(w, constval.StatusCodeServerError, err.Error(), "", response)
		return
	}
}

// FormUpload receive file, check and save to disk
func FormUpload(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var response = make(map[string]interface{})

	// set max upload size
	r.Body = http.MaxBytesReader(w, r.Body, constval.MaxFileUploadSize)

	// set max buffer size
	err := r.ParseMultipartForm(constval.MaxFileUploadSize)
	if err != nil {
		utils.SetReponse(w, constval.StatusCodeForbidden, err.Error(), "", response)
		return
	}

	// get file handler
	file, handler, err := r.FormFile("file")
	if err != nil {
		utils.SetReponse(w, constval.StatusCodeForbidden, err.Error(), "", response)
		return
	}
	// check if file is image
	if !strings.HasPrefix(utils.GetFileMIMEType(filepath.Ext(handler.Filename)), "image/") {
		utils.SetReponse(w, constval.StatusCodeForbidden, "", "Upload only image", response)
		return
	}
	// check if auth key is valid
	if !strings.EqualFold(r.FormValue("auth"), CONFIG.AuthFormKey) {
		utils.SetReponse(w, constval.StatusCodeForbidden, "", "Auth key is not correct", response)
		return
	}

	// parse and save file
	if file != nil {
		defer file.Close()
		fileName, err := utils.SaveToDisk(constval.ImageFolderPath, file, filepath.Ext(handler.Filename))
		if err != nil {
			utils.SetReponse(w, constval.StatusCodeForbidden, err.Error(), "", response)
			return
		}

		response["file"] = fileName
	}

	utils.SetReponse(w, constval.StatusCodeOk, "", "", response)
}
