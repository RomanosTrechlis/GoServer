package admin

import (
	"github.com/RomanosTrechlis/GoServer/server/util"
	"net/http"
	"github.com/RomanosTrechlis/GoServer/server/logger"
)

func LoginGet(w http.ResponseWriter, r *http.Request) {
	util.Templates.ExecuteTemplate(w, "login.html", nil);
}

func LoginPost(w http.ResponseWriter, r *http.Request) {
	session, _ := Store.Get(r, "cookie-name")

	// Authentication goes here
	r.ParseForm()
	user := User{}
	user.Username = r.PostForm.Get("username")
	user.Password = r.PostForm.Get("password")
	logger.Debug.Println("username: ", user.Username)
	password, _ := HashPassword("password")

	if !CheckPasswordHash(user.Password, password) {
		util.Templates.ExecuteTemplate(w, "login.html", Error{ErrorMessage: "Please enter the correct username and password."});
		return
	}

	// Set user as authenticated
	session.Values["authenticated"] = true
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusFound)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := Store.Get(r, "cookie-name")

	// Revoke users authentication
	session.Values["authenticated"] = false
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusFound)
}
