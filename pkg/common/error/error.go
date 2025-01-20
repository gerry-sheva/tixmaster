package error

import (
	"net/http"

	"github.com/gerry-sheva/tixmaster/pkg/util"
)

func ErrorResponse(w http.ResponseWriter, status int, message any) {
	envelope := util.Envelope{"error": message}

	err := util.NewRwJSON().Write(w, status, envelope, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func ServerErrorResponse(w http.ResponseWriter) {
	msg := "The server encountered a problem and could not process your request"
	ErrorResponse(w, http.StatusInternalServerError, msg)
}

func NotFoundResponse(w http.ResponseWriter) {
	msg := "The requested resource could not be found"
	ErrorResponse(w, http.StatusNotFound, msg)
}

func MethodNotAllowedResponse(w http.ResponseWriter) {
	msg := "Unsupported method"
	ErrorResponse(w, http.StatusMethodNotAllowed, msg)
}

func FailedValidationResponse(w http.ResponseWriter, errors map[string]string) {
	ErrorResponse(w, http.StatusUnprocessableEntity, errors)
}
