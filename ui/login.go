package main

import (
	"fmt"
	"log"
	"net/http"
)

func (app *App) login(w http.ResponseWriter, r *http.Request) {
	login := r.FormValue("login")
	password := r.FormValue("password")

	sql_result := app.DB.QueryRow("select * from user_data where login = $1", login)

	var user User_Data
	if err := sql_result.Scan(&user.ID_User, &user.Login, &user.Password); err != nil || user.Login == "" {
		err_message := "User with such login doesn't exists"
		fmt.Fprintf(w, err_message)
		log.Println(err_message)
		return
	}

	if user.Password != password {
		err_message := "Login or password don't match"
		fmt.Fprintf(w, err_message)
		log.Printf(err_message)
		return
	}

	cookie := &http.Cookie{
		Name:  "authorized",
		Value: login,
		Path:  "/",
	}

	http.SetCookie(w, cookie)
	w.Header().Add("HX-Redirect", "/protected/home.html")
}
