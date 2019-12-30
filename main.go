package main

import (
	"database/sql"
	"fmt"
	"net/url"
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
	_, err := sql.Open(`mysql`, dsn)
	if err != nil {
		fmt.Println(err)
	}
}