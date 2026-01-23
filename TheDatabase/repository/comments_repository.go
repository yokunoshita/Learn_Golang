package repository

import (
	"TheDatabase/entity"
	"context"
)

type CommentRepository interface {
	Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error)
	FindById(ctx context.Context, id int32) (entity.Comment, error)
	FindMany(ctx context.Context) ([]entity.Comment, error)
}
