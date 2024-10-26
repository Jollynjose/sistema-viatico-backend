package middlewares

import "net/http"

type Middleware func(http.HandlerFunc) http.HandlerFunc

func Chain(middlewares ...Middleware) Middleware {
	return func(handler http.HandlerFunc) http.HandlerFunc {
		for _, middleware := range middlewares {
			handler = middleware(handler)
		}
		return handler
	}
}
