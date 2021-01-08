package middleware

import (
	ServiceAuth "api/services"
	"api/util"
	"net/http"
)

// Auth -
func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		err := ServiceAuth.TokenValid(header)

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(util.MessageError(err))
			return
		}

		next.ServeHTTP(w, r)
	}
}
