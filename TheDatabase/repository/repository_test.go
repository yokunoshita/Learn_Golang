package repository

import (
	TheDatabase "TheDatabase"
	"TheDatabase/entity"
	"context"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentsAdd(t *testing.T) {
	db := TheDatabase.GetConnection()
	defer db.Close()
	commentRepo := NewCommentsRepo(db)

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "testrepo@go.dev",
		Comment: "Added by repo",
	}
	result, err := commentRepo.Insert(ctx, comment)

	if err != nil {
		t.Fatal("Failed to add comment : ", err)
	}
	fmt.Println("Comment added with ID : ", result.Id)
}

func TestCommentsFindId(t *testing.T) {
	commentRepo := NewCommentsRepo(TheDatabase.GetConnection())

	comment, err := commentRepo.FindById(context.Background(), 64)
	if err != nil {
		t.Fatal("Failed to look up the id")
	}
	fmt.Println(comment)
}

func TestCommentsFindAll(t *testing.T) {
	commentRepo := NewCommentsRepo(TheDatabase.GetConnection())

	comments, err := commentRepo.FindMany(context.Background())
	if err != nil {
		t.Fatal("Failed to look up the data")
	}

	for _, comment := range comments {
		fmt.Println(comment)
	}

}
