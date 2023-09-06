package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	Port     string
	Handlers map[string]http.HandlerFunc
	Router   chi.Router
}

func NewServer(port string) *Server {
	return &Server{
		Port:     port,
		Handlers: make(map[string]http.HandlerFunc),
		Router:   chi.NewRouter(),
	}
}

func (s *Server) AddHandler(path string, handler http.HandlerFunc) {
	s.Handlers[path] = handler
}

func (s *Server) Start() {
	s.Router.Use(middleware.Logger)
	for path, handler := range s.Handlers {
		s.Router.Handle(path, handler)
	}

	srv := &http.Server{
		Addr:    s.Port,
		Handler: s.Router,
	}

	go func() {
		log.Println("Listening and serving on port: ", s.Port)
		if err := srv.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	signal.Stop(quit)
	close(quit)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	log.Println("Shutting down server...")
	if err := srv.Shutdown(ctx); err != nil {
		panic(err)
	}
	log.Println("Server successfully ended!")

}
