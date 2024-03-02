package response

import "net/http"

type Responser interface {
	EncodeResponse(w http.ResponseWriter)
}

type Err struct {
	Code int
}

func (e Err) EncodeResponse(w http.ResponseWriter) {
	w.WriteHeader(e.Code)
}

type Nope struct{}

func (n Nope) EncodeResponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusTeapot)
}

type Ok struct{}

func (o Ok) EncodeResponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
}
