package belajar_golang_database

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/belajar-golang")
	if err != nil {
		panic(err)
	}

	defer db.Close()
}

func TestLagi(t *testing.T) {
	SayTest()
}
