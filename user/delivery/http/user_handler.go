package http

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/gustavonahum/code-validator/user"
)

type UserHandler struct {
	UUsecase user.Usecase
}

func NewUserHandler(r *mux.Router, u user.Usecase) {
	handler := &UserHandler{
		UUsecase: u,
	}
	s := r.PathPrefix("/user").Subrouter()
	s.HandleFunc("/{id}", handler.GetById)
}

func (h UserHandler) GetById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("olar\n"))
}
