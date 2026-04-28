package http

import (
	"context"
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

	return &HTTPServer{router: mux.NewRouter(), httpHandlers: handlerHTTP}, nil
}

func (s *HTTPServer) StartServer() {
	s.router.Path("/post").Methods(http.MethodPost).HandlerFunc(s.httpHandlers.CreatePost) // created Post
}
