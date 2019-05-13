package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"

	"github.com/eunkenlow/golang-api-boilerplate/pkg/server/middlewares"
	"github.com/eunkenlow/golang-api-boilerplate/pkg/server/routes"
)

// routers returns all routes
func routers() http.Handler {
	r := chi.NewRouter()
	// router middlewares
	r.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.RequestID,
		middleware.RealIP,
		middleware.Logger,
		middleware.Recoverer,
		middleware.Timeout(60*time.Second),
	)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Golang api boilerplate"))
	})
	// define routes
	r.Route("/v1", func(r chi.Router) {
		r.Use(middlewares.IsAuthorized)
		r.Mount("/me", routes.MeRouter())
		r.Mount("/users", routes.UserRouter())
	})

	return r
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Println("Server running on port", port)

	router := routers()
	log.Fatal(http.ListenAndServe(":"+port, router))
}
