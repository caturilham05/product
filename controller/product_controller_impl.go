package controller

import (
	"caturilham05/product/helper"
	"caturilham05/product/model/web"
	"caturilham05/product/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type ProductControllerImpl struct {
	ProductService service.ProductService
}

// ClaimToken implements ProductController.
func (p *ProductControllerImpl) ClaimToken(writer http.ResponseWriter, request *http.Request, params httprouter.Params) helper.JWTClaims {
	accessToken := request.Header.Get("Authorization")
	if accessToken == "" {
		return helper.JWTClaims{}
	}

	claims, err := helper.ValidateToken(accessToken)
	if err != nil {
		return helper.JWTClaims{}
	}

	return *claims
}

// Create implements ProductController.
func (p *ProductControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productCreateRequest := web.ProductCreateRequest{}
	helper.ReadFromRequestBody(request, &productCreateRequest)

	productResponse := p.ProductService.Create(request.Context(), productCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   productResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// Delete implements ProductController.
func (p *ProductControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	claims := p.ClaimToken(writer, request, params)
	id, err := strconv.Atoi(params.ByName("productId"))
	helper.PanicIfError(err)

	p.ProductService.Delete(request.Context(), id, claims.Id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   nil,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// FindAll implements ProductController.
func (p *ProductControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productResponses := p.ProductService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   productResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// FindById implements ProductController.
func (p *ProductControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("productId"))
	helper.PanicIfError(err)

	productResponse := p.ProductService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   productResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// Update implements ProductController.
func (p *ProductControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productUpdateRequest := web.ProductUpdateRequest{}
	helper.ReadFromRequestBody(request, &productUpdateRequest)
	id, err := strconv.Atoi(params.ByName("productId"))
	helper.PanicIfError(err)

	productUpdateRequest.Id = id

	productResponse := p.ProductService.Update(request.Context(), productUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   productResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func NewProductController(productService service.ProductService) ProductController {
	return &ProductControllerImpl{
		ProductService: productService,
	}
}
