package http

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/gustavonahum/code-validator/models"
	"github.com/gustavonahum/code-validator/user"
)

type UserHandler struct {
	uUsecase user.Usecase
}

func NewUserHandler(r *mux.Router, u user.Usecase) {

	handler := &UserHandler{
		uUsecase: u,
	}

	r.HandleFunc("/user", handler.Store).Methods("POST")
	r.HandleFunc("/user/{id}", handler.GetById).Methods("GET")
}

func (h *UserHandler) GetById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		log.Printf("Error reading path variable: %v", err)
		http.Error(w, "can't convert path variable to int64", http.StatusBadRequest)
		return
	}

	user, _ := h.uUsecase.GetById(id)
	userBytesSlice, err := json.Marshal(user)
	if err != nil {
		log.Printf("Error converting struct to byte slice: %v", user)
		http.Error(w, "can't convert struct to byte slice", http.StatusInternalServerError)
		return
	}

	w.Write(userBytesSlice)
}

func (h *UserHandler) Store(w http.ResponseWriter, r *http.Request) {

	var user models.User

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Printf("Error unmarshalling body: %v", err)
		http.Error(w, "can't unmarshal body", http.StatusUnprocessableEntity)
		return
	}

	h.uUsecase.Store(&user)
}
