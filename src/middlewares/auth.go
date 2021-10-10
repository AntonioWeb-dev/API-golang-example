package middlewares

import (
	"api/src/authentication"
	"api/src/utils"
	"log"
	"net/http"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

func Authentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := authentication.VerifyToken(r); err != nil {
			utils.Error(w, http.StatusUnauthorized, err)
			return
		}
		next(w, r)
	}
}
