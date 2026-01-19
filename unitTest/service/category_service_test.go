package service

import (
	"testing"
	"unitTest/entity"
	"unitTest/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepo = &repository.CategoryRepoMoc{Mock: mock.Mock{}}
var categoryServ = CategoryService{Repository: categoryRepo}

func TestCategoryService_Get(t *testing.T) {
	categoryRepo.Mock.On("FindById", "1").Return(nil)
	category, err := categoryServ.Get("1")
	assert.NotNil(t, err)
	assert.Nil(t, category)
}

func TestCategoryService_GetFound(t *testing.T) {
	category := entity.Category{
		Id:   "2",
		Name: "Handphone",
	}
	categoryRepo.Mock.On("FindById", "3").Return(category)

	res, err := categoryServ.Get("2")
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, category.Id, res.Id)
	assert.Equal(t, category.Name, res.Name)

}
