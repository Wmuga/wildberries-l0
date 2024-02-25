package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type StatusRecorder struct {
	http.ResponseWriter
	Status int
}

func (r *StatusRecorder) WriteHeader(status int) {
	r.Status = status
	r.ResponseWriter.WriteHeader(status)
}

func GetRequestLogger(logger *log.Logger) mux.MiddlewareFunc {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			recorder := &StatusRecorder{w, 200}
			start := time.Now()
			h.ServeHTTP(recorder, r)
			dur := time.Since(start)
			log.Println(r.URL, recorder.Status, dur)
		})
	}
}
