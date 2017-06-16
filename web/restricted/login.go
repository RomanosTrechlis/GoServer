package restricted

import (
	"net/http"

	"github.com/RomanosTrechlis/GoServer/logger"
	c "github.com/RomanosTrechlis/GoServer/util/conf"
)

func LoginGet(w http.ResponseWriter, r *http.Request) {
	c.Templates.ExecuteTemplate(w, "login.html", nil)
}

func LoginPost(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	// Authentication goes here
	r.ParseForm()
	user := user{}
	user.Username = r.PostForm.Get("username")
	user.Password = r.PostForm.Get("password")
	logger.Debug.Println("username: ", user.Username)
	password, _ := hashPassword("password")

	if !checkPasswordHash(user.Password, password) {
		c.Templates.ExecuteTemplate(w, "login.html", Error{ErrorMessage: "Please enter the correct username and password."})
		return
	}

	// Set user as authenticated
	session.Values["authenticated"] = true
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusFound)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	// Revoke users authentication
	session.Values["authenticated"] = false
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusFound)
}
