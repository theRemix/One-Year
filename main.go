package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/alexandrevicenzi/go-sse"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

func main() {
	events = sse.NewServer(nil)
	defer events.Shutdown()
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: true,
	}))

	workDir, _ := os.Getwd()
	staticDir := http.Dir(filepath.Join(workDir, "client/public"))
	FileServer(r, "/", staticDir)

	r.Mount("/api", apiRouter())
	r.Mount("/events", events)
	r.Get("/join", spaHandler)
	r.Get("/player", spaHandler)
	r.Get("/host", spaHandler)
	r.Get("/stage", spaHandler)
	r.Get("/create", spaHandler)

	log.Fatal("HTTP server error: ", http.ListenAndServe(":4000", r))
}

// catch all for spa routes
func spaHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Add("Content Type", "text/html")
	fmt.Println(req.URL.Path)
	http.ServeFile(rw, req, "./client/public/index.html")
}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
