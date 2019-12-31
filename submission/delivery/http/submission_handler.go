package http

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/gustavonahum/code-validator/models"
	"github.com/gustavonahum/code-validator/submission"
)

type SubmissionHandler struct {
	sUsecase submission.Usecase
}

func NewSubmissionHandler(r *mux.Router, s submission.Usecase) {

	handler := &SubmissionHandler{
		sUsecase: s,
	}

	r.HandleFunc("/user/{id}", handler.GetById).Methods("GET")
	r.HandleFunc("/user", handler.Store).Methods("POST")
	r.HandleFunc("/user/{id}", handler.Delete).Methods("DELETE")
}

func (h *SubmissionHandler) GetById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		log.Printf("Error reading path variable: %v", err)
		http.Error(w, "can't convert path variable to int64", http.StatusBadRequest)
		return
	}

	submission, _ := h.sUsecase.GetById(id)
	submissionBytesSlice, err := json.Marshal(submission)
	if err != nil {
		log.Printf("Error converting struct to byte slice: %v", submission)
		http.Error(w, "can't convert struct to byte slice", http.StatusUnprocessableEntity)
		return
	}

	w.Write(submissionBytesSlice)
}

func (h *SubmissionHandler) Store(w http.ResponseWriter, r *http.Request) {

	var submission models.Submission

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &submission)
	if err != nil {
		log.Printf("Error unmarshalling body: %v", err)
		http.Error(w, "can't unmarshal body", http.StatusUnprocessableEntity)
		return
	}

	h.sUsecase.Store(&submission)
}

func (h *SubmissionHandler) Delete(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		log.Printf("Error reading path variable: %v", err)
		http.Error(w, "can't convert path variable to int64", http.StatusBadRequest)
		return
	}

	submission, _ := h.sUsecase.Delete(id)
	submissionBytesSlice, err := json.Marshal(submission)
	if err != nil {
		log.Printf("Error converting struct to byte slice: %v", submission)
		http.Error(w, "can't convert struct to byte slice", http.StatusUnprocessableEntity)
		return
	}

	w.Write(submissionBytesSlice)
}
