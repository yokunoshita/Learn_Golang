package repository

import "unitTest/entity"

type CategoryRepo interface {
	FindById(id string) *entity.Category
}
