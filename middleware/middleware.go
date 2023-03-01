package middleware

import (
	"log"
	"net/http"
	"time"
)

// LoggingHandler
// User it like this:
// http.Handle(("/url"), LoggingHandler(http.HandlerFunc(do)))
func LoggingHanlder(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Started %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Printf("Completed %s in %v", r.URL.Path, time.Since(start))
	})
}
