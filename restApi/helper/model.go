package helper

import (
	"restApi/model/domain"
	"restApi/model/web"
)

func ToCategoryRes(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var categoryResponse []web.CategoryResponse
	for _, category := range categories {
		categoryResponse = append(categoryResponse, ToCategoryRes(category))
	}

	return categoryResponse
}
