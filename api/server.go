package api

import (
	"fmt"
	"net/http"

	config "upload-form/config"

	"github.com/gorilla/mux"
)

type muxR struct {
	r *mux.Router
}

// StartServer - start server using mux
func StartServer() {
	// ec2 router
	fmt.Println(http.ListenAndServe(":"+config.Port, &muxR{LoadRouter()}))
}

func (s *muxR) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	// cors configuration
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Methods", "GET,OPTIONS,POST,PUT")
	res.Header().Set("Access-Control-Allow-Headers", "Content-Type,X-Amz-Date,X-Amz-Security-Token,Authorization,X-Api-Key,X-Requested-With,Accept,Access-Control-Allow-Methods,Access-Control-Allow-Origin,Access-Control-Allow-Headers")

	if req.Method == "OPTIONS" {
		return
	}

	s.r.ServeHTTP(res, req)
}
