package main

import (
	"fmt"
	"log"
	"net/http"
)

func (app *App) signup(w http.ResponseWriter, r *http.Request) {
	login := r.FormValue("login")
	password := r.FormValue("password")

	_, err := app.DB.Exec("insert into user_data (login, password) values ($1, $2)", login, password)
	if err != nil {
		fmt.Fprintln(w, err)
		log.Println(err)
		return
	}

	cookie := &http.Cookie{
		Name:  "authorized",
		Value: "true",
		Path:  "/",
	}

	http.SetCookie(w, cookie)
	w.Header().Add("HX-Redirect", "/home.html")
}
