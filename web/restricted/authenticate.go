package restricted

import (
	"net/http"

	"github.com/RomanosTrechlis/GoServer/util/middleware"
	"github.com/gorilla/sessions"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

type Error struct {
	ErrorMessage string
}

func Authenticate() middleware.Middleware {
	// Create a new Middleware
	return func(f http.HandlerFunc) http.HandlerFunc {
		// Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {
			// Do middleware things
			session, _ := store.Get(r, "cookie-name")
			// Check if user is authenticated
			if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
				//http.Error(w, "Forbidden", http.StatusForbidden)
				http.Redirect(w, r, "/login/", http.StatusFound)
				return
			}
			f(w, r)
		}
	}
}
