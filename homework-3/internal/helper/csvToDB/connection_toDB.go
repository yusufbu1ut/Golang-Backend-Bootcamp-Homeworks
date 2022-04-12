package csvToDB

import (
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-yusufbu1ut/internal/infrastructure"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-yusufbu1ut/internal/models/author"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-yusufbu1ut/internal/models/book"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-yusufbu1ut/internal/models/bookAuthor"
)

func ToConnectDB() (*book.BookRepository, *author.AuthorRepository, *bookAuthor.BookAuthRepository) {
	var bookRepository *book.BookRepository
	var authorRepository *author.AuthorRepository
	var bookAuthorRepository *bookAuthor.BookAuthRepository
	db := infrastructure.NewPostgresDB("host=localhost user=postgres password=pass1234 dbname=library port=5432 sslmode=disable") // arrange it for your db connections
	bookRepository = book.NewBookRepository(db)
	authorRepository = author.NewAuthorRepository(db)
	bookAuthorRepository = bookAuthor.NewBookAuthRepository(db)
	bookRepository.Migration()
	authorRepository.Migration()
	bookAuthorRepository.Migration()
	bookRepository, authorRepository, bookAuthorRepository = ReadBookWithWorkerPool("../docs/books2.csv", bookRepository, authorRepository, bookAuthorRepository) //Take comment after fist run

	return bookRepository, authorRepository, bookAuthorRepository
}
