package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type ApplicationController struct{}

func (ac *ApplicationController) Write(w http.ResponseWriter, data interface{}) {
	result, err := json.Marshal(data)

	if err != nil {
		ac.Error(w, ServerError{err})
	} else {
		fmt.Fprint(w, string(result))
	}
}

func (ac *ApplicationController) Error(w http.ResponseWriter, err error) {

	code := 400
	if err.Error() == "not found" {
		code = 404
	} else if _, ok := err.(ServerError); ok {
		code = 500
	}

	w.WriteHeader(code)
	fmt.Fprintf(w, "{\"error\": \"%s\"}", err)
}

func (ac *ApplicationController) NotFound(w http.ResponseWriter) {
	ac.Error(w, errors.New("not found"))
}

type ServerError struct {
	error
}