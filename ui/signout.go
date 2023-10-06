package main

import "net/http"

func (app *App) signout(w http.ResponseWriter, r *http.Request) {
  cookie := &http.Cookie {
		Name:  "authorized",
		Value: "",
		Path:  "/",
  }

  http.SetCookie(w, cookie)
	w.Header().Add("HX-Redirect", "/")
}
