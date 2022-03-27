package util

import (
	"encoding/json"
	"net/http"
	"strconv"
	constval "upload-form/constant"
)

// set request response with status, message
func SetReponse(w http.ResponseWriter, status, errorMsg, msg string, response map[string]interface{}) {
	w.Header().Set("Status", status)
	response["meta"] = setMeta(status, errorMsg, msg)
	s, _ := strconv.Atoi(status)
	w.WriteHeader(s)
	json.NewEncoder(w).Encode(response)
}

func setMeta(status, errorMsg, msg string) map[string]string {
	if len(msg) == 0 {
		if status == constval.StatusCodeBadRequest {
			msg = "Bad Request"
		} else if status == constval.StatusCodeServerError {
			msg = "Server Error"
		}
	}
	return map[string]string{
		"status":        status,
		"error_message": errorMsg,
		"message":       msg,
	}
}

func getHTTPStatusCode(code string) int {
	switch code {
	case constval.StatusCodeOk:
		return http.StatusOK
	case constval.StatusCodeCreated:
		return http.StatusCreated
	case constval.StatusCodeBadRequest:
		return http.StatusBadRequest
	case constval.StatusCodeServerError:
		return http.StatusInternalServerError
	}
	return http.StatusOK
}
