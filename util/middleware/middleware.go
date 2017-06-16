package middleware

import (
	"net/http"
	"time"

	"github.com/RomanosTrechlis/GoServer/logger"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

// Chain applies middlewares to a http.HandlerFunc
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
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
