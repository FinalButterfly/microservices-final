package app

import (
	"net/http"
	"strings"

	u "../utils"
)

var JwtAuthentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		notAuth := []string{"/api/user/new", "/api/user/login"}
		requestpath := r.URL.Path

		for _, value := range notAuth {
			if value == requestpath {
				next.ServeHTTP(w, r)
				return
			}
		}

		tokenHeader := r.Header.Get("Authorization")
		if tokenHeader == "" {
			negativeResponse("Missing auth token", w)
			return
		}

		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) != 2 {
			negativeResponse("Invalid/Malformed auth token", w)
			return
		}

		tokenPart := splitted[1]

		if tokenPart != "123" {
			negativeResponse("Invalid API token", w)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func negativeResponse(message string, w http.ResponseWriter) {
	response := make(map[string]interface{})
	response = u.Message(false, "Invalid/Malformed auth token")
	w.WriteHeader(http.StatusForbidden)
	w.Header().Add("Content-Type", "application/json")
	u.Respond(w, response)
	return
}
