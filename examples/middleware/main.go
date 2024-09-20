package main

import (
	"context"
	"fmt"
	"net/http"

	tss "github.com/lucarin91/go-typesafe-serve"
)

type Source string

func Logger(next tss.HandlerFunc) tss.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s\n", r.Method, r.URL.Path)

		next(w, r)
	}
}

func CheckSource(next tss.HandlerFunc1[Source]) tss.HandlerFunc {
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

func WhoCtrl(_ context.Context, req tss.Nope, source Source) (tss.Ok, *tss.Err) {
	if source != "me" {
		panic("this will never happen, source will be always 'me'")
	}

	return tss.Ok{}, nil
}

func main() {
	r := http.NewServeMux()

	middlewares := tss.Compose2(Logger, CheckSource)

	r.HandleFunc("GET /who", middlewares(tss.ToHandler1(WhoCtrl)))

	fmt.Println("Listening on :8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
