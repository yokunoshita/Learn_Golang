package service

import (
	"context"
	"database/sql"
	"restApi/helper"
	"restApi/model/domain"
	"restApi/model/web"
	"restApi/repository"

	"github.com/go-playground/validator/v10"
)

type CategoryServiceImp struct {
	CategoryRepository repository.CategoryRepo
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCategoryService(categoryRepo repository.CategoryRepo, DB *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImp{
		CategoryRepository: categoryRepo,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *CategoryServiceImp) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name: request.Name,
	}
	category = service.CategoryRepository.Create(ctx, tx, category)
	return helper.ToCategoryRes(category)
}

func (service *CategoryServiceImp) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	category.Name = request.Name

	category = service.CategoryRepository.Update(ctx, tx, category)
	return helper.ToCategoryRes(category)
}

func (service *CategoryServiceImp) Delete(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	helper.PanicIfError(err)

	service.CategoryRepository.Delete(ctx, tx, category)
}

func (service *CategoryServiceImp) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	helper.PanicIfError(err)

	return helper.ToCategoryRes(category)
}

func (service *CategoryServiceImp) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := service.CategoryRepository.FindAll(ctx, tx)
	return helper.ToCategoryResponses(categories)
}
