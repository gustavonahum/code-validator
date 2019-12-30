package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/labstack/echo"

	//_studentCodeHttpDeliver "github.com/gustavonahum/code-validator/studentCode/delivery/http"
	_studentCodeRepo "github.com/gustavonahum/code-validator/studentCode/repository"
	_studentCodeUcase "github.com/gustavonahum/code-validator/studentCode/usecase"
)

func main() {
	dbHost := "localhost"
	dbPort := "5601"
	dbUser := "user"
	dbPass := "pass"
	dbName := "name"
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "America/Sao Paulo")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := sql.Open(`mysql`, dsn)

	e := echo.New()
	studentCodeRepo := _studentCodeRepo.NewMysqlStudentCodeRepository(dbConn)
	ar := _studentCodeRepo.NewMysqlStudentCodeRepository(dbConn)

	timeoutContext := time.Duration(10000) * time.Second
	au := _studentCodeUcase.NewStudentCodeUsecase(ar, studentCodeRepo, timeoutContext)
	//_studentCodeHttpDeliver.NewStudentCodeHandler(e, au)
}