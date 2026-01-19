package service

import (
	"errors"
	"unitTest/entity"
	"unitTest/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepo
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)
	if category == nil {
		return category, errors.New("Category not found")
	} else {
		return category, nil
	}
}
