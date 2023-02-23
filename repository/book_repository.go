package repository

import (
	"context"
	"go-rest-without-framework/model"
)

type BookRepository interface {
	Save(ctx context.Context, book model.Book)
	Update(ctx context.Context, book model.Book)
	Delete(ctx context.Context, book model.Book)
	FindById(ctx context.Context, bookId int) (model.Book, error)
	FindAll(ctx context.Context) []model.Book
}
