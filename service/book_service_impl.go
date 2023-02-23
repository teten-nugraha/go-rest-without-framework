package service

import (
	"context"
	"go-rest-without-framework/data/request"
	"go-rest-without-framework/data/response"
	"go-rest-without-framework/helpers"
	"go-rest-without-framework/model"
	"go-rest-without-framework/repository"
)

type BookServiceImpl struct {
	BookRepository repository.BookRepository
}

func (b *BookServiceImpl) Create(ctx context.Context, request request.BookCreateRequest) {
	book := model.Book{
		Name: request.Name,
	}

	b.BookRepository.Save(ctx, book)
}

func (b *BookServiceImpl) Update(ctx context.Context, request request.BookUpdateRequest) {
	book, err := b.BookRepository.FindById(ctx, request.Id)
	helpers.PanicIfError(err)

	book.Name = request.Name
	b.BookRepository.Update(ctx, book)
}

func (b *BookServiceImpl) Delete(ctx context.Context, bookId int) {
	book, err := b.BookRepository.FindById(ctx, bookId)
	helpers.PanicIfError(err)
	b.BookRepository.Delete(ctx, book)
}

func (b *BookServiceImpl) FindById(ctx context.Context, bookId int) response.BookResponse {
	book, err := b.BookRepository.FindById(ctx, bookId)
	helpers.PanicIfError(err)
	return response.BookResponse(book)
}

func (b *BookServiceImpl) FindAll(ctx context.Context) []response.BookResponse {
	books := b.BookRepository.FindAll(ctx)

	var bookResp []response.BookResponse

	for _, value := range books {
		book := response.BookResponse{Id: value.Id, Name: value.Name}
		bookResp = append(bookResp, book)
	}

	return bookResp
}

func NewBookServiceImpl(bookRepository repository.BookRepository) BookService {
	return &BookServiceImpl{BookRepository: bookRepository}
}
