package belajar_golang_database

import (
	"context"
	"database/sql"
	"fmt"
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
