package tss

import (
	"context"
	"encoding/json"
	"net/http"
)

func decode[T Requester[T]](w http.ResponseWriter, r *http.Request) (T, bool) {
	var req T
	req, err := req.DecodeRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		// TODO: make this configurable
		_ = json.NewEncoder(w).Encode(struct {
			Err string `json:"err"`
		}{Err: err.Error()})
		return req, false
	}
	return req, true
}

func encode[U Responser, V Responser](w http.ResponseWriter, res U, errRes *V) {
	if errRes != nil {
		(*errRes).EncodeResponse(w)
		return
	}
	res.EncodeResponse(w)
}

type Func[T Requester[T], U Responser, V Responser] func(ctx context.Context, req T) (U, *V)

func ToHandler[T Requester[T], U Responser, V Responser](ctr Func[T, U, V]) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req, ok := decode[T](w, r)
		if !ok {
			return
		}

		res, errRes := ctr(r.Context(), req)

		encode(w, res, errRes)
	}
}

type Func1[T Requester[T], U Responser, V Responser, A any] func(ctx context.Context, req T, a A) (U, *V)

func ToHandler1[T Requester[T], U Responser, V Responser, A any](ctr Func1[T, U, V, A]) HandlerFunc1[A] {
	return func(w http.ResponseWriter, r *http.Request, a A) {
		req, ok := decode[T](w, r)
		if !ok {
			return
		}

		res, errRes := ctr(r.Context(), req, a)

		encode(w, res, errRes)
	}
}

type Func2[T Requester[T], U Responser, V Responser, A any, B any] func(ctx context.Context, req T, a A, b B) (U, *V)

func ToHandler2[T Requester[T], U Responser, V Responser, A any, B any](ctr Func2[T, U, V, A, B]) HandlerFunc2[A, B] {
	return func(w http.ResponseWriter, r *http.Request, a A, b B) {
		req, ok := decode[T](w, r)
		if !ok {
			return
		}

		res, errRes := ctr(r.Context(), req, a, b)

		encode(w, res, errRes)
	}
}

type Func3[T Requester[T], U Responser, V Responser, A any, B any, C any] func(ctx context.Context, req T, a A, b B, c C) (U, *V)

func ToHandler3[T Requester[T], U Responser, V Responser, A any, B any, C any](ctr Func3[T, U, V, A, B, C]) HandlerFunc3[A, B, C] {
	return func(w http.ResponseWriter, r *http.Request, a A, b B, c C) {
		req, ok := decode[T](w, r)
		if !ok {
			return
		}

		res, errRes := ctr(r.Context(), req, a, b, c)

		encode(w, res, errRes)
	}
}

type Func4[T Requester[T], U Responser, V Responser, A any, B any, C any, D any] func(ctx context.Context, req T, a A, b B, c C, d D) (U, *V)

func ToHandler4[T Requester[T], U Responser, V Responser, A any, B any, C any, D any](ctr Func4[T, U, V, A, B, C, D]) HandlerFunc4[A, B, C, D] {
	return func(w http.ResponseWriter, r *http.Request, a A, b B, c C, d D) {
		req, ok := decode[T](w, r)
		if !ok {
			return
		}

		res, errRes := ctr(r.Context(), req, a, b, c, d)

		encode(w, res, errRes)
	}
}

type Func5[T Requester[T], U Responser, V Responser, A any, B any, C any, D any, E any] func(ctx context.Context, req T, a A, b B, c C, d D, e E) (U, *V)

func ToHandler5[T Requester[T], U Responser, V Responser, A any, B any, C any, D any, E any](ctr Func5[T, U, V, A, B, C, D, E]) HandlerFunc5[A, B, C, D, E] {
	return func(w http.ResponseWriter, r *http.Request, a A, b B, c C, d D, e E) {
		req, ok := decode[T](w, r)
		if !ok {
			return
		}

		res, errRes := ctr(r.Context(), req, a, b, c, d, e)

		encode(w, res, errRes)
	}
}

type Func6[T Requester[T], U Responser, V Responser, A any, B any, C any, D any, E any, F any] func(ctx context.Context, req T, a A, b B, c C, d D, e E, f F) (U, *V)

func ToHandler6[T Requester[T], U Responser, V Responser, A any, B any, C any, D any, E any, F any](ctr Func6[T, U, V, A, B, C, D, E, F]) HandlerFunc6[A, B, C, D, E, F] {
	return func(w http.ResponseWriter, r *http.Request, a A, b B, c C, d D, e E, f F) {
		req, ok := decode[T](w, r)
		if !ok {
			return
		}

		res, errRes := ctr(r.Context(), req, a, b, c, d, e, f)

		encode(w, res, errRes)
	}
}

type Func7[T Requester[T], U Responser, V Responser, A any, B any, C any, D any, E any, F any, G any] func(ctx context.Context, req T, a A, b B, c C, d D, e E, f F, g G) (U, *V)

func ToHandler7[T Requester[T], U Responser, V Responser, A any, B any, C any, D any, E any, F any, G any](ctr Func7[T, U, V, A, B, C, D, E, F, G]) HandlerFunc7[A, B, C, D, E, F, G] {
	return func(w http.ResponseWriter, r *http.Request, a A, b B, c C, d D, e E, f F, g G) {
		req, ok := decode[T](w, r)
		if !ok {
			return
		}

		res, errRes := ctr(r.Context(), req, a, b, c, d, e, f, g)

		encode(w, res, errRes)
	}
}

type Func8[T Requester[T], U Responser, V Responser, A any, B any, C any, D any, E any, F any, G any, H any] func(ctx context.Context, req T, a A, b B, c C, d D, e E, f F, g G, h H) (U, *V)

func ToHandler8[T Requester[T], U Responser, V Responser, A any, B any, C any, D any, E any, F any, G any, H any](ctr Func8[T, U, V, A, B, C, D, E, F, G, H]) HandlerFunc8[A, B, C, D, E, F, G, H] {
	return func(w http.ResponseWriter, r *http.Request, a A, b B, c C, d D, e E, f F, g G, h H) {
		req, ok := decode[T](w, r)
		if !ok {
			return
		}

		res, errRes := ctr(r.Context(), req, a, b, c, d, e, f, g, h)

		encode(w, res, errRes)
	}
}

type Func9[T Requester[T], U Responser, V Responser, A any, B any, C any, D any, E any, F any, G any, H any, I any] func(ctx context.Context, req T, a A, b B, c C, d D, e E, f F, g G, h H, i I) (U, *V)

func ToHandler9[T Requester[T], U Responser, V Responser, A any, B any, C any, D any, E any, F any, G any, H any, I any](ctr Func9[T, U, V, A, B, C, D, E, F, G, H, I]) HandlerFunc9[A, B, C, D, E, F, G, H, I] {
	return func(w http.ResponseWriter, r *http.Request, a A, b B, c C, d D, e E, f F, g G, h H, i I) {
		req, ok := decode[T](w, r)
		if !ok {
			return
		}

		res, errRes := ctr(r.Context(), req, a, b, c, d, e, f, g, h, i)

		encode(w, res, errRes)
	}
}

type Func10[T Requester[T], U Responser, V Responser, A any, B any, C any, D any, E any, F any, G any, H any, I any, J any] func(ctx context.Context, req T, a A, b B, c C, d D, e E, f F, g G, h H, i I, j J) (U, *V)

func ToHandler10[T Requester[T], U Responser, V Responser, A any, B any, C any, D any, E any, F any, G any, H any, I any, J any](ctr Func10[T, U, V, A, B, C, D, E, F, G, H, I, J]) HandlerFunc10[A, B, C, D, E, F, G, H, I, J] {
	return func(w http.ResponseWriter, r *http.Request, a A, b B, c C, d D, e E, f F, g G, h H, i I, j J) {
		req, ok := decode[T](w, r)
		if !ok {
			return
		}

		res, errRes := ctr(r.Context(), req, a, b, c, d, e, f, g, h, i, j)

		encode(w, res, errRes)
	}
}
