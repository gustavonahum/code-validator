package main

import (
	"database/sql"

	"github.com/gorilla/mux"
	_userRepo "github.com/gustavonahum/code-validator/user/repository"
	_userCodeRepo "github.com/gustavonahum/code-validator/userCode/repository"

	_userUsecase "github.com/gustavonahum/code-validator/user/usecase"
	_userCodeUsecase "github.com/gustavonahum/code-validator/userCode/usecase"

	_userHandler "github.com/gustavonahum/code-validator/user/delivery/http"
)

func main() {

	var dbConn *sql.DB

	userRepo := _userRepo.NewPostgresqlUserRepository(dbConn)
	userCodeRepo := _userCodeRepo.NewPostgresqlUserCodeRepository(dbConn)

	userUsecase := _userUsecase.NewUserUsecase(userRepo)
	_userCodeUsecase.NewUserCodeUsecase(userCodeRepo, userRepo)

	r := mux.NewRouter()
	_userHandler.NewUserHandler(r, userUsecase)
}
