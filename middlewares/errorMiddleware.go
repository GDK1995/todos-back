package middlewares

import (
	"errors"
	"log"
	"net/http"
)

func ErrorMiddleware(handler func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := handler(w, r); err != nil {
			var httpErr *HTTPError
			if errors.As(err, &httpErr) {
				http.Error(w, httpErr.Message, httpErr.Code)
				return
			}
			log.Println(err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
	}
}
