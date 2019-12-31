package main

import (
	"database/sql"

	_userRepo "github.com/gustavonahum/code-validator/user/repository"
	_userCodeRepo "github.com/gustavonahum/code-validator/userCode/repository"

	_userCodeUsecase "github.com/gustavonahum/code-validator/userCode/usecase"
)

func main() {
	var dbConn *sql.DB

	userRepo := _userRepo.NewPostgresqlUserRepository(dbConn)
	userCodeRepo := _userCodeRepo.NewPostgresqlUserCodeRepository(dbConn)

	_userCodeUsecase.NewUserCodeUsecase(userCodeRepo, userRepo)

}
