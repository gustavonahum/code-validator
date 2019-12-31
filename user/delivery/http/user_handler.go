package http

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/gustavonahum/code-validator/user"
)

type UserHandler struct {
	uUsecase user.Usecase
}

func NewUserHandler(r *mux.Router, u user.Usecase) {
	handler := &UserHandler{
		uUsecase: u,
	}
	s := r.PathPrefix("/user").Subrouter()
	s.HandleFunc("/{id}", handler.GetById).Methods("GET")
	s.HandleFunc("/", handler.Store).Methods("POST")
}

func (h *UserHandler) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	user, _ := h.uUsecase.GetById(id)
	userByteArr, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(userByteArr)
}

func (h *UserHandler) Store(w http.ResponseWriter, r *http.Request) {

}
