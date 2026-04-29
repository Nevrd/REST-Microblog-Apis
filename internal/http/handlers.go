package http

import (
	"API/internal/database"
	"API/internal/errors"
	"API/internal/model"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
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
		return
	}
	err = h.db.InsertPost(post)
	if err != nil {
		errModel := model.NewErrorModel(err)
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(errModel); err != nil {
			fmt.Println(errors.ErrFromEncodeJson)
		}
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(model.ValidateEncodePostModel(post))
}

func (h *HTTPHandlers) GetPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := h.db.GetAllPost()
	if err != nil {
		errModel := model.NewErrorModel(err)
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(errModel); err != nil {
			fmt.Println(errors.ErrFromEncodeJson)
		}
		return
	}
	w.WriteHeader(http.StatusOK)
	for _, v := range posts {
		json.NewEncoder(w).Encode(v)
	}
}

func (h *HTTPHandlers) GetPost(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]
	post, err := h.db.GetPost(title)
	if err != nil {
		errModel := model.NewErrorModel(err)
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(errModel); err != nil {
			fmt.Println(errors.ErrFromEncodeJson)
		}
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)
}

func (h *HTTPHandlers) UpdatePost(w http.ResponseWriter, r *http.Request) {
	post := model.PostModel{}
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		errModel := model.NewErrorModel(err)
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(errModel); err != nil {
			fmt.Println(errors.ErrFromEncodeJson)
		}
		return
	}
	err = h.db.UpdatePost(post)
	if err != nil {
		errModel := model.NewErrorModel(err)
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(errModel); err != nil {
			fmt.Println(errors.ErrFromEncodeJson)
		}
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *HTTPHandlers) DeletePost(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]
	err := h.db.DeletePost(title)
	if err != nil {
		errModel := model.NewErrorModel(err)
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(errModel); err != nil {
			fmt.Println(errors.ErrFromEncodeJson)
		}
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
