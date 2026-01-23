package repository

import (
	"TheDatabase/entity"
	"context"
	"database/sql"
	"errors"
	"strconv"
)

type commentsRepoImpl struct {
	DB *sql.DB
}

func NewCommentsRepo(db *sql.DB) CommentRepository {
	return &commentsRepoImpl{DB: db}
}

func (repo *commentsRepoImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	query := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	result, err := repo.DB.ExecContext(ctx, query, comment.Email, comment.Comment)
	if err != nil {
		return comment, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return comment, err
	}

	comment.Id = int(id)
	return comment, nil
}

func (repo *commentsRepoImpl) FindById(ctx context.Context, id int32) (entity.Comment, error) {
	query := "SELECT id, email, comment FROM comments WHERE id = ? LIMIT 1"
	comment := entity.Comment{}
	err := repo.DB.QueryRowContext(ctx, query, id).Scan(
		&comment.Id,
		&comment.Email,
		&comment.Comment,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return comment, errors.New("Comment with ID " + strconv.Itoa(int(id)) + " not found")
		}
		return comment, err
	}
	return comment, nil
}

func (repo *commentsRepoImpl) FindMany(ctx context.Context) ([]entity.Comment, error) {
	query := "SELECT id, email, comment FROM comments"
	rows, err := repo.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []entity.Comment
	for rows.Next() {
		comment := entity.Comment{}
		rows.Scan(
			&comment.Id,
			&comment.Email,
			&comment.Comment,
		)
		comments = append(comments, comment)
	}
	return comments, nil
}
