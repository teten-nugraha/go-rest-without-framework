package repository

import (
	"context"
	"database/sql"
	"errors"
	"go-rest-without-framework/helpers"
	"go-rest-without-framework/model"
)

type BookRepositoryImpl struct {
	Db *sql.DB
}

func NewBookRepositoryImpl(Db *sql.DB) BookRepository {
	return &BookRepositoryImpl{Db: Db}
}

func (b *BookRepositoryImpl) Save(ctx context.Context, book model.Book) {
	tx, err := b.Db.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	SQL := "insert into book(name) values($1)"
	_, err = tx.ExecContext(ctx, SQL, book.Name)
	helpers.PanicIfError(err)
}

func (b *BookRepositoryImpl) Update(ctx context.Context, book model.Book) {
	tx, err := b.Db.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	SQL := "update book set name=$1 where id =$2"
	_, err = tx.ExecContext(ctx, SQL, book.Name, book.Id)
	helpers.PanicIfError(err)
}

func (b *BookRepositoryImpl) Delete(ctx context.Context, book model.Book) {
	tx, err := b.Db.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	SQL := "delete from book where id = $1"
	_, err = tx.ExecContext(ctx, SQL, book.Id)
	helpers.PanicIfError(err)
}

func (b *BookRepositoryImpl) FindById(ctx context.Context, bookId int) (model.Book, error) {
	tx, err := b.Db.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	SQL := "select id,name from book where id = $1"
	result, err := tx.QueryContext(ctx, SQL, bookId)
	helpers.PanicIfError(err)
	defer result.Close()

	book := model.Book{}

	if result.Next() {
		err := result.Scan(&book.Id, &book.Name)
		helpers.PanicIfError(err)
		return book, nil
	} else {
		return book, errors.New("book id not found")
	}

}

func (b *BookRepositoryImpl) FindAll(ctx context.Context) []model.Book {
	tx, err := b.Db.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	SQL := "select id, name from book"
	result, err := tx.QueryContext(ctx, SQL)
	defer result.Close()

	var books []model.Book

	for result.Next() {
		book := model.Book{}
		err := result.Scan(&book.Id, &book.Name)
		helpers.PanicIfError(err)

		books = append(books, book)
	}

	return books
}
