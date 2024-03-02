package request

import "net/http"

type Requester[T any] interface {
	DecodeRequest(r *http.Request) (T, error)
}

type Nope struct{}

func (n Nope) DecodeRequest(r *http.Request) (Nope, error) {
	return Nope{}, nil
}
