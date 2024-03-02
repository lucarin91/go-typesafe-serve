package tsserve

import "net/http"

type HandlerFunc = http.HandlerFunc
type HandlerFunc1[A any] func(w http.ResponseWriter, r *http.Request, a A)
type HandlerFunc2[A any, B any] func(w http.ResponseWriter, r *http.Request, a A, b B)
type HandlerFunc3[A any, B any, C any] func(w http.ResponseWriter, r *http.Request, a A, b B, c C)
type HandlerFunc4[A any, B any, C any, D any] func(w http.ResponseWriter, r *http.Request, a A, b B, c C, d D)
type HandlerFunc5[A any, B any, C any, D any, E any] func(w http.ResponseWriter, r *http.Request, a A, b B, c C, d D, e E)
type HandlerFunc6[A any, B any, C any, D any, E any, F any] func(w http.ResponseWriter, r *http.Request, a A, b B, c C, d D, e E, f F)
type HandlerFunc7[A any, B any, C any, D any, E any, F any, G any] func(w http.ResponseWriter, r *http.Request, a A, b B, c C, d D, e E, f F, g G)
type HandlerFunc8[A any, B any, C any, D any, E any, F any, G any, H any] func(w http.ResponseWriter, r *http.Request, a A, b B, c C, d D, e E, f F, g G, h H)
type HandlerFunc9[A any, B any, C any, D any, E any, F any, G any, H any, I any] func(w http.ResponseWriter, r *http.Request, a A, b B, c C, d D, e E, f F, g G, h H, i I)
type HandlerFunc10[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any] func(w http.ResponseWriter, r *http.Request, a A, b B, c C, d D, e E, f F, g G, h H, i I, j J)
