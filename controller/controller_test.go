package controller

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	data string
}

func (Input) DecodeRequest(r *http.Request) (Input, error) {
	return Input{
		data: r.URL.Query().Get("data"),
	}, nil
}

type Output struct {
	data []byte
}

func (o Output) EncodeResponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(o.data)
	if err != nil {
		panic(err)
	}
}

type ErrorOutput struct {
	data []byte
}

func (e ErrorOutput) EncodeResponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	_, err := w.Write(e.data)
	if err != nil {
		panic(err)
	}
}

func TestController(t *testing.T) {
	tests := []struct {
		name           string
		req            *http.Request
		output         Output
		errorOutput    *ErrorOutput
		expectedInput  Input
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "ok",
			req:  httptest.NewRequest(http.MethodGet, "/?data=hello", nil),
			output: Output{
				data: []byte("Hello, world!"),
			},
			expectedInput: Input{
				data: "hello",
			},
			expectedStatus: http.StatusOK,
			expectedBody:   "Hello, world!",
		},
		{
			name: "error",
			req:  httptest.NewRequest(http.MethodGet, "/?data=hello", nil),
			errorOutput: &ErrorOutput{
				data: []byte("Error!!"),
			},
			expectedInput:  Input{data: "hello"},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "Error!!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()

			controller := func(_ context.Context, input Input) (Output, *ErrorOutput) {
				assert.Equal(t, tt.expectedInput, input)
				return tt.output, tt.errorOutput
			}
			handler := http.HandlerFunc(ToHandler(controller))

			handler.ServeHTTP(rr, tt.req)
			assert.Equal(t, tt.expectedStatus, rr.Code)
			assert.Equal(t, tt.expectedBody, rr.Body.String())
		})
	}
}

func TestDecode(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/?data=hello", nil)

	out, ok := decode[Input](w, r)
	assert.True(t, ok)
	assert.Equal(t, Input{data: "hello"}, out)

	// check no error
	assert.Equal(t, w.Code, 200)
	assert.Equal(t, w.Body.String(), "")
}

func TestEncode(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		w := httptest.NewRecorder()

		res := Output{data: []byte("Hello, world!")}
		encode[Output, ErrorOutput](w, res, nil)

		assert.Equal(t, w.Code, 200)
		assert.Equal(t, w.Body.String(), "Hello, world!")
	})

	t.Run("error", func(t *testing.T) {
		w := httptest.NewRecorder()

		err := ErrorOutput{data: []byte("Error!!")}
		encode(w, Output{}, &err)

		assert.Equal(t, w.Code, 500)
		assert.Equal(t, w.Body.String(), "Error!!")
	})
}
