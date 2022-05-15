package middleware

import (
	"log"
	"net/http"
)

type statusCaptureResponseWriter struct {
	http.ResponseWriter
	status      int
	wroteHeader bool
}

func wrapResponseWriter(w http.ResponseWriter) *statusCaptureResponseWriter {
	return &statusCaptureResponseWriter{ResponseWriter: w}
}

func (rw *statusCaptureResponseWriter) Status() int {
	return rw.status
}

func (rw *statusCaptureResponseWriter) WriteHeader(code int) {
	if rw.wroteHeader {
		return
	}
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
	rw.wroteHeader = true
	return
}

func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wrappedRW := wrapResponseWriter(w)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(wrappedRW, r)

		// Do stuff here
		log.Println(r.Method, r.RequestURI, wrappedRW.status)

	})
}
