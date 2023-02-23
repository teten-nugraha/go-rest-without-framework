package controller

import (
	"github.com/julienschmidt/httprouter"
	"go-rest-without-framework/data/request"
	"go-rest-without-framework/data/response"
	"go-rest-without-framework/helpers"
	"go-rest-without-framework/service"
	"net/http"
	"strconv"
)

type BookController struct {
	BookService service.BookService
}

func NewBookController(bookService service.BookService) *BookController {
	return &BookController{BookService: bookService}
}

func (c *BookController) Create(w http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	bookCreateRequest := request.BookCreateRequest{}
	helpers.ReadPayload(requests, &bookCreateRequest)

	c.BookService.Create(requests.Context(), bookCreateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}
	helpers.WriteResponseBody(w, webResponse)
}

func (c *BookController) Update(w http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	bookUpdateRequest := request.BookUpdateRequest{}
	helpers.ReadPayload(requests, &bookUpdateRequest)

	bookId := params.ByName("bookId")
	id, err := strconv.Atoi(bookId)
	helpers.PanicIfError(err)

	bookUpdateRequest.Id = id

	c.BookService.Update(requests.Context(), bookUpdateRequest)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	helpers.WriteResponseBody(w, webResponse)
}

func (controller *BookController) Delete(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	bookId := params.ByName("bookId")
	id, err := strconv.Atoi(bookId)
	helpers.PanicIfError(err)

	controller.BookService.Delete(requests.Context(), id)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	helpers.WriteResponseBody(writer, webResponse)

}

func (controller *BookController) FindAll(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	result := controller.BookService.FindAll(requests.Context())
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	helpers.WriteResponseBody(writer, webResponse)
}

func (controller *BookController) FindById(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	bookId := params.ByName("bookId")
	id, err := strconv.Atoi(bookId)
	helpers.PanicIfError(err)

	result := controller.BookService.FindById(requests.Context(), id)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	helpers.WriteResponseBody(writer, webResponse)

}
