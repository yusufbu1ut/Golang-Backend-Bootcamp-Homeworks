package csvToDB

import (
	"encoding/csv"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-yusufbu1ut/internal/models/author"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-yusufbu1ut/internal/models/book"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-yusufbu1ut/internal/models/bookAuthor"
)

var (
	repbook     *book.BookRepository
	repauthor   *author.AuthorRepository
	repbookauth *bookAuthor.BookAuthRepository
)

type item struct {
	bks   book.Book
	auths string
}

var Cnt = 0

func ReadBookWithWorkerPool(path string, trepbook *book.BookRepository, trepauthor *author.AuthorRepository, trepbookauth *bookAuthor.BookAuthRepository) (*book.BookRepository, *author.AuthorRepository, *bookAuthor.BookAuthRepository) {
	repbook = trepbook
	repauthor = trepauthor
	repbookauth = trepbookauth
	linesChan := make(chan []string)
	resultsChan := make(chan item)
	wgB := new(sync.WaitGroup)
	for w := 1; w <= 3; w++ {
		wgB.Add(1)
		go convertToItemsStruct(linesChan, resultsChan, wgB)
	}

	go func() {
		//reading lines
		f, _ := os.Open(path)
		defer f.Close()
		lines, _ := csv.NewReader(f).ReadAll()
		isFirstRow := true
		for _, line := range lines {
			if isFirstRow {
				isFirstRow = false
				continue
			}

			linesChan <- line
		}
		close(linesChan)
	}()

	go func() {
		wgB.Wait()

		close(resultsChan)
	}()

	for v := range resultsChan {
		repbook.Create(v.bks)
		sepAddAuthors(v.auths, v.bks)
	}
	return repbook, repauthor, repbookauth
}

// convertToItemsStruct seperates lines to items
func convertToItemsStruct(booksChan <-chan []string, resultschan chan<- item, wg *sync.WaitGroup) {
	defer wg.Done()

	for b := range booksChan {
		isbn, _ := strconv.Atoi(b[5])
		_book := book.NewBook(b[1], isbn, b[7], b[11])
		_item := item{
			bks:   *_book,
			auths: b[2],
		}
		resultschan <- _item
	}
}

//sepAddAuthors in here creating new bookAuthor elements for book_authors db
func sepAddAuthors(auths string, book book.Book) {
	authsRes := strings.Split(auths, "/")
	for _, a := range authsRes {
		_author := isInAuthors(a)
		_bookAuthor := bookAuthor.NewBook_Author(book.ISBN, _author.ID)
		repbookauth.Create(*_bookAuthor)
		Cnt++
	}
}

//isInAuthors checks is in allread authors if it is not adds to authors
func isInAuthors(authr string) author.Author {
	_author := author.NewAuth(authr)

	repauthor.Create(*_author)
	tauthor, _ := repauthor.FindByAuthor(*_author)

	return tauthor
}
