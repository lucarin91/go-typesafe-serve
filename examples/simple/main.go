package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	tss "github.com/lucarin91/go-typesafe-serve"
)

type HelloReq struct {
	Name string
}

func (r HelloReq) DecodeRequest(req *http.Request) (HelloReq, error) {
	return HelloReq{
		Name: req.PathValue("name"),
	}, nil
}

type HelloRes struct {
	Message string `json:"message"`
}

func (r HelloRes) EncodeResponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(r)
}

func HelloCtrl(_ context.Context, req HelloReq) (HelloRes, *tss.Err) {
	return HelloRes{Message: "Hello, " + req.Name + "!"}, nil
}

func main() {
	r := http.NewServeMux()
	r.HandleFunc("GET /hello/{name}", tss.ToHandler(HelloCtrl))

	fmt.Println("Listening on :8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
