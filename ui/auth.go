package main

import (
	"log"
	"net/http"
)

func requireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth, err := r.Cookie("authorized")
		if err != nil || auth.Value != "true" {
      log.Println("Unauthorized access to protected content:", r.URL.Path)
      http.Redirect(w, r, "/index.html", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}
