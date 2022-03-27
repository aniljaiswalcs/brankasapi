package config

import "os"

var AuthFormKey string
var Port string

func LoadConfig() {

	AuthFormKey = os.Getenv("AUTH_FORM_KEY")
	Port = os.Getenv("PORT")

}
