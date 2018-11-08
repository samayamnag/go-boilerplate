package handlers

import (
	"encoding/json"
	"net/http"
)

// Source: https://github.com/w3tecch/go-api-boilerplate/blob/master/app/lib/Response.go

type Response struct {
	ResponseWriter	http.ResponseWriter
}

func (r *Response) SendOK(body interface{}) {
	setJSON(r.ResponseWriter)
	encodeJSON(r.ResponseWriter, body)
}

func (r *Response) SendCreated(body interface{}) {
	setJSON(r.ResponseWriter)
	setHttpStatus(r.ResponseWriter, http.StatusCreated)
	encodeJSON(r.ResponseWriter, body)
}

func (r *Response) SendBadRequest(message string) {
	http.Error(r.ResponseWriter, message, http.StatusBadRequest)
}

func (r *Response) SendUnauthenticated(message string) {
	http.Error(r.ResponseWriter, message, http.StatusUnauthorized)
}

func (r *Response) SendForbidden(message string) {
	http.Error(r.ResponseWriter, message, http.StatusForbidden)
}

func (r *Response) SendNotFound(message string) {
	http.Error(r.ResponseWriter, message, http.StatusNotFound)
}

func (r *Response) SendValidationErrors(errors []byte) {
	setJSON(r.ResponseWriter)
	setHttpStatus(r.ResponseWriter, http.StatusUnprocessableEntity)
	r.ResponseWriter.Write(errors)
}

func (r *Response) SendUnprocessableEntity(errors map[string]interface{}) {
	setJSON(r.ResponseWriter)
	setHttpStatus(r.ResponseWriter, http.StatusUnprocessableEntity)
	encodeJSON(r.ResponseWriter, errors)
}

func (r *Response) SendServerError(message string) {
	http.Error(r.ResponseWriter, message, http.StatusInternalServerError)
}

// Local fuctions
func setJSON(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func setHttpStatus(w http.ResponseWriter, status int)  {
	w.WriteHeader(status)
}

func encodeJSON(w http.ResponseWriter, body interface{}) {
	json.NewEncoder(w).Encode(body)
}

