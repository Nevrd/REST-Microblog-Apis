package http

import (
	"API/internal/database"
	"API/internal/errors"
	"API/internal/model"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type HTTPHandlers struct {
	db *database.Database
}

func NewHTTPHandlers(ctx context.Context) (*HTTPHandlers, error) {
	path := "postgres://postgres:0404@172.25.96.1:5432/postgres"
	db, err := database.NewDataBase(ctx, path)

	if err != nil {
		return &HTTPHandlers{}, err
	}

	return &HTTPHandlers{db: db}, nil
}

func (h *HTTPHandlers) CreatePost(w http.ResponseWriter, r *http.Request) {
	post := model.PostModel{}
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		errModel := model.NewErrorModel(err)
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(errModel); err != nil {
			fmt.Println(errors.ErrFromEncodeJson)
		}
	}

}
