package belajar_golang_database

import (
	"context"
	"fmt"
	"strconv"
	"testing"
)

func TestDatabaseTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	// do transaction
	script := "INSERT INTO  comments(email, comment) VALUES(?, ?)"
	statement, err := tx.PrepareContext(ctx, script)
	if err != nil {
		panic(err)
	}

	defer statement.Close()

	for i := 0; i < 10; i++ {
		email := "Pengguna " + strconv.Itoa(i) + "@gmail.com"
		comment := "Comment ke " + strconv.Itoa(i)

		result, err := statement.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Comment ID", id)
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		panic(err)
	}

}
