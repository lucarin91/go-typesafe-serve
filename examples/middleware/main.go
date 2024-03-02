package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lucarin91/go-typesafe-serve/controller"
	"github.com/lucarin91/go-typesafe-serve/request"
	"github.com/lucarin91/go-typesafe-serve/response"
	"github.com/lucarin91/go-typesafe-serve/tsserve"
)

type Source string

func Logger(next tsserve.HandlerFunc) tsserve.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s\n", r.Method, r.URL.Path)

		next(w, r)
	}
}

func CheckSource(next tsserve.HandlerFunc1[Source]) tsserve.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		source := r.Header.Get("source")

		switch source {
		case "":
			w.WriteHeader(http.StatusBadRequest)
		case "me":
			next(w, r, Source(source))
		default:
			w.WriteHeader(http.StatusForbidden)
			return
		}
	}
}

func WhoCtrl(_ context.Context, req request.Nope, source Source) (response.Ok, *response.Err) {
	if source != "me" {
		panic("this will never happen, source will be always 'me'")
	}

	return response.Ok{}, nil
}

func main() {
	r := http.NewServeMux()

	middlewares := tsserve.Compose2(Logger, CheckSource)

	r.HandleFunc("GET /who", middlewares(controller.ToHandler1(WhoCtrl)))

	fmt.Println("Listening on :8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
