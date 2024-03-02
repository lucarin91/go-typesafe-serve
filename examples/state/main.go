package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lucarin91/go-typesafe-serve/controller"
	"github.com/lucarin91/go-typesafe-serve/response"
)

type MyState map[string]string

type NumReq struct {
	Num string
}

func (r NumReq) DecodeRequest(req *http.Request) (NumReq, error) {
	return NumReq{
		Num: req.PathValue("num"),
	}, nil
}

type NumRes struct {
	Num string `json:"num"`
}

func (r NumRes) EncodeResponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(r)
}

func NumCtrl(state MyState) controller.Func[NumReq, NumRes, response.Err] {
	return func(_ context.Context, req NumReq) (NumRes, *response.Err) {
		numStr, ok := state[req.Num]
		if !ok {
			return NumRes{}, &response.Err{
				Code: http.StatusNotFound,
			}
		}

		return NumRes{Num: numStr}, nil
	}
}

func main() {
	state := MyState{
		"1": "one",
		"2": "two",
		"3": "three",
	}

	r := http.NewServeMux()
	r.HandleFunc("GET /num/{num}", controller.ToHandler(NumCtrl(state)))

	fmt.Println("Listening on :8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
