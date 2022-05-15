package middleware

import (
	"github.com/gorilla/mux"
	"net/http"
)

func GlobalHeaderInjector(headers map[string]string) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Do stuff here
			for k, v := range headers {
				w.Header().Set(k, v)
			}
			// Call the next handler, which can be another middleware in the chain, or the final handler.
			next.ServeHTTP(w, r)

		})
	}
}
