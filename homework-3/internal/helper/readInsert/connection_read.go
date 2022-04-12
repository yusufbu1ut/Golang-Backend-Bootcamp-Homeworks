package readInsert

import (
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-yusufbu1ut/internal/infrastructure"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-yusufbu1ut/internal/models/author"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-yusufbu1ut/internal/models/book"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-yusufbu1ut/internal/models/bookAuthor"
)

func ReadConnectDB() (*bookAuthor.BookAuthRepository, *author.AuthorRepository, *book.BookRepository) {
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
	books, auths, bookAuths, _ := ReadBookWithWorkerPool("../docs/books2.csv") //Take comment after fist run
	bookRepository.InsertSampleData(books)                                     //Take comment after fist run
	authorRepository.InsertSampleData(auths)                                   //Take comment after fist run
	bookAuthorRepository.InsertSampleData(bookAuths)                           //Take comment after fist run

	return bookAuthorRepository, authorRepository, bookRepository
}
