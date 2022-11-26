package belajar_golang_database

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "INSERT INTO customer(id, name, email, balance, rating, birth_date, married) VALUES('','Test','test@gmail.com','80000','8.5','1999-10-10', false)"
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name FROM customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}

		fmt.Println("id = ", id)
		fmt.Println("name = ", name)
	}
}

func TestQuerySelect(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var birthDate sql.NullTime
		var createdAt time.Time
		var married bool

		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}

		fmt.Println("Id: ", id)
		fmt.Println("Name: ", name)
		fmt.Println("Email: ", email)
		fmt.Println("Balance: ", balance)
		fmt.Println("Rating: ", rating)
		fmt.Println("BirthDate: ", birthDate)
		fmt.Println("Married: ", married)
		fmt.Println("CreatedAt: ", createdAt)
	}
}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	username := "admin'; #"
	password := "admin"

	// Not Secure
	// script := "SELECT username FROM users WHERE username = '" + username + "' AND password = '" + password + "'"

	// Secure
	script := "SELECT username FROM users WHERE username = ? AND password = ?"
	rows, err := db.QueryContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Hallo ", username)
	} else {
		fmt.Println("Gagal")
	}
}

func TestGetLastInserted(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	email := "eko@gmail.com"
	comment := "comments"

	ctx := context.Background()
	script := "INSERT INTO comments(email, comment) VALUES(?,?)"
	result, err := db.ExecContext(ctx, script, email, comment)
	if err != nil {
		panic(err)
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("Success With ID = ", insertId)
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	statement, err := db.PrepareContext(ctx, script)
	if err != nil {
		panic(err)
	}

	defer statement.Close()

	for i := 1; i <= 10; i++ {
		email := "User" + strconv.Itoa(i) + "@gmail.com"
		comment := "Komentar ke " + strconv.Itoa(i)
		result, err := statement.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("Comment ID ", id)
	}
}
