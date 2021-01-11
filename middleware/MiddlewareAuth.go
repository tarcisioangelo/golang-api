package middleware

import (
	ServiceAuth "api/services"
	"api/util"
	"net/http"

	"github.com/gorilla/context"
)

// Auth -
func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		userID, err := ServiceAuth.TokenValid(header)

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(util.MessageError(err))
			return
		}

		context.Set(r, "userID", userID)

		next.ServeHTTP(w, r)
	}
}
