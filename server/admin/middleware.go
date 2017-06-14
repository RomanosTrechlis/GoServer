package admin

import (
	"net/http"
	"github.com/gorilla/sessions"
	"time"
	"github.com/RomanosTrechlis/GoServer/server/logger"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key = []byte("super-secret-key")
	Store = sessions.NewCookieStore(key)
)

type User struct {
	Username string
	Password string
}

type Error struct {
	ErrorMessage string
}

type Middleware func(http.HandlerFunc) http.HandlerFunc

func Authenticate() Middleware {
	// Create a new Middleware
	return func(f http.HandlerFunc) http.HandlerFunc {
		// Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {
			logger.Info.Println("Auth")
			// Do middleware things
			session, _ := Store.Get(r, "cookie-name")
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

func Logging() Middleware {
	// Create a new Middleware
	return func(f http.HandlerFunc) http.HandlerFunc {
		// Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {
			// Do middleware things
			start := time.Now()
			defer func() { logger.Info.Println(r.URL.Path, time.Since(start)) }()

			// Call the next middleware/handler in chain
			f(w, r)
		}
	}
}

// Chain applies middlewares to a http.HandlerFunc
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}
