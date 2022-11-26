package test

import (
	belajar_golang_database "belajar-golang-database"
	"belajar-golang-database/entity"
	repository "belajar-golang-database/repository"
	"context"
	"fmt"
	"testing"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := repository.NewCommentRepository(belajar_golang_database.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "repository@gmail.com",
		Comment: "Test Repository 2",
	}

	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	commentRepository := repository.NewCommentRepository(belajar_golang_database.GetConnection())

	comment, err := commentRepository.FindById(context.Background(), 100)
	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
}

func TestFindAll(t *testing.T) {
	commentRepository := repository.NewCommentRepository(belajar_golang_database.GetConnection())

	comments, err := commentRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}

	for _, comment := range comments {
		fmt.Println(comment)
	}
}
