package tss

import "net/http"

type Requester[T any] interface {
	DecodeRequest(r *http.Request) (T, error)
}

func (n Nope) DecodeRequest(r *http.Request) (Nope, error) {
	return Nope{}, nil
}
