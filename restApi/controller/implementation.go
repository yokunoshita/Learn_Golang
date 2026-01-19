package controller

import (
	"net/http"
	"restApi/helper"
	"restApi/model/web"
	"restApi/service"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CatetegoryControllerImp struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CatetegoryControllerImp{
		CategoryService: categoryService,
	}
}

func (controller *CatetegoryControllerImp) Create(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	var categoryCreateRequest web.CategoryCreateRequest
	helper.ReadFromReqBody(request, &categoryCreateRequest)

	categoryResponse := controller.CategoryService.Create(request.Context(), categoryCreateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusCreated,
		Status: "Ok",
		Data:   categoryResponse,
	}

	helper.WriteToRes(writer, webResponse)
}

func (controller *CatetegoryControllerImp) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryUpdateRequest := web.CategoryUpdateRequest{}
	helper.ReadFromReqBody(request, &categoryUpdateRequest)

	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryUpdateRequest.Id = id

	categoryResponse := controller.CategoryService.Update(request.Context(), categoryUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   categoryResponse,
	}

	helper.WriteToRes(writer, webResponse)
}

func (controller *CatetegoryControllerImp) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	controller.CategoryService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
	}

	helper.WriteToRes(writer, webResponse)
}

func (controller *CatetegoryControllerImp) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryResponse := controller.CategoryService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   categoryResponse,
	}

	helper.WriteToRes(writer, webResponse)
}

func (controller *CatetegoryControllerImp) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryResponse := controller.CategoryService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   categoryResponse,
	}

	helper.WriteToRes(writer, webResponse)
}
