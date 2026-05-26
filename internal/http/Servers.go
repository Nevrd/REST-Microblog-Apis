package http

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPServer struct {
	router       *mux.Router
	httpHandlers *HTTPHandlers
}

// context(ctx) for database(waitCancel)
func NewServer(ctx context.Context) (*HTTPServer, error) {

	handlerHTTP, err := NewHTTPHandlers(ctx)
	if err != nil {
		return &HTTPServer{}, err
	}
	return &HTTPServer{router: mux.NewRouter(), httpHandlers: handlerHTTP}, err
}

func (s *HTTPServer) StartServer() error {
	s.router.Path("/posts").Methods(http.MethodPost).HandlerFunc(s.httpHandlers.CreatePost) // created Post
	s.router.Path("/posts").Methods(http.MethodGet).HandlerFunc(s.httpHandlers.GetPosts)
	s.router.Path("/posts/{title}").Methods(http.MethodGet).HandlerFunc(s.httpHandlers.GetPost)
	s.router.Path("/posts/{title}").Methods(http.MethodDelete).HandlerFunc(s.httpHandlers.DeletePost)

	log.Println("Server Started")
	err := http.ListenAndServe(":8080", s.router)
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}
