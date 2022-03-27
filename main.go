package main

import (
	formapi "upload-form/api"
	config "upload-form/config"
)

func main() {

	config.LoadConfig()
	formapi.StartServer()
}
