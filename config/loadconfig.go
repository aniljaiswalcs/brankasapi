package config

var AuthFormKey string
var Port string

func LoadConfig() {

	//AuthFormKey = os.Getenv("AUTH_FORM_KEY")
	AuthFormKey = "jennny"
	//Port = os.Getenv("PORT")
	Port = "8000"

}
