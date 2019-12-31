package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_submissionRepo "github.com/gustavonahum/code-validator/submission/repository"
	_userRepo "github.com/gustavonahum/code-validator/user/repository"

	_submissionUsecase "github.com/gustavonahum/code-validator/submission/usecase"
	_userUsecase "github.com/gustavonahum/code-validator/user/usecase"

	_userHandler "github.com/gustavonahum/code-validator/user/delivery/http"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {

	var dbConn *sql.DB

	userRepo := _userRepo.NewPostgresqlUserRepository(dbConn)
	submissionRepo := _submissionRepo.NewPostgresqlSubmissionRepository(dbConn)

	userUsecase := _userUsecase.NewUserUsecase(userRepo)
	_submissionUsecase.NewSubmissionUsecase(submissionRepo)

	r := mux.NewRouter()
	_userHandler.NewUserHandler(r, userUsecase)

	http.ListenAndServe(":8080", r)
}
