package middleware

import (
	"log"
	"net/http"
)

type AuthenticationMiddleware struct {
	tokenUsers map[string]string
}

func (am *AuthenticationMiddleware) Populate() {
	am.tokenUsers = make(map[string]string)
	am.tokenUsers["000000000"] = "guest"
	am.tokenUsers["111111111"] = "userB"
	am.tokenUsers["222222222"] = "userC"
	am.tokenUsers["333333333"] = "userD"
}

func (am *AuthenticationMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Session-Token")

		if user, found := am.tokenUsers[token]; found {
			log.Printf("Authenticated user %s\n", user)
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}
