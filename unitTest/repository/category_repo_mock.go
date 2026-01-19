package repository

import (
	"unitTest/entity"

	"github.com/stretchr/testify/mock"
)

type CategoryRepoMoc struct {
	Mock mock.Mock
}

func (repository *CategoryRepoMoc) FindById(id string) *entity.Category {
	args := repository.Mock.Called(id)
	if args.Get(0) == nil {
		return nil
	}

	category := args.Get(0).(entity.Category)
	return &category
}
