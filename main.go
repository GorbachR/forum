package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path"
	"strings"
	"syscall"

	"github.com/GorbachR/forum/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	listenPort := flag.String("listenport", ":3000", "the port the server is listening on")

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("Error getting workdin dir:", err)
	}
	filesDir := http.Dir(path.Join(wd, "dist/assets"))
	FileServer(r, "/assets", filesDir)

	r.Get("/", handlers.Root)
	r.Get("/login", handlers.Login)

	server := http.Server{
		Handler: r,
		Addr:    *listenPort,
	}

	go func() {
		server.ListenAndServe()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatal("Server shutdown:", err)
	}

	log.Print("Server exiting")

}

func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.Contains(path, "{}*") {
		log.Fatal("Fileserver does not permit any URL parameters")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", http.StatusMovedPermanently).ServeHTTP)
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
