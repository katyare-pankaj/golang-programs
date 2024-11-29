// Dependency imports
package main

import (
	"net/http"

	"github.com/justinas/alice"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Middleware chain to handle healthcheck requests
	healthCheckMiddleware := alice.New(HealthCheckMiddleware).Then(e.Router())

	// Healthcheck endpoint
	e.GET("/healthcheck", echo.WrapHandler(healthCheckMiddleware))

	// Other API routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":1323"))
}

// HealthCheckMiddleware handles healthcheck requests
func HealthCheckMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/healthcheck" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
			return
		}
		h.ServeHTTP(w, r)
	})
}
