package readInsert

import (
	"encoding/csv"
	"strconv"

	"os"
	"strings"
	"sync"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-yusufbu1ut/internal/models/author"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-yusufbu1ut/internal/models/book"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-yusufbu1ut/internal/models/bookAuthor"
)

//item strct is for seperating and holding elemnts when workerpool starts
type item struct {
	bks   book.Book
	auths string
}

func ReadBookWithWorkerPool(path string) ([]book.Book, []author.Author, []bookAuthor.Book_Author, error) {
	var ResultsBooks []book.Book
	var ResultsAuthors []author.Author
	var ResultsBookAuth []bookAuthor.Book_Author

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
	//	var results []item
	// for i := range resultsChan {

	// 	results = append(results, i)
	// }
	for v := range resultsChan {
		ResultsBooks = append(ResultsBooks, v.bks) //Books appending to add to db.books
		ResultsBookAuth, ResultsAuthors = sepAddAuthors(v.auths, v.bks, ResultsBookAuth, ResultsAuthors)
	}

	return ResultsBooks, ResultsAuthors, ResultsBookAuth, nil
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
func sepAddAuthors(auths string, book book.Book, ResultsBookAuth []bookAuthor.Book_Author, ResultsAuthors []author.Author) ([]bookAuthor.Book_Author, []author.Author) {
	authsRes := strings.Split(auths, "/")
	for _, a := range authsRes {
		var _author author.Author
		_author, ResultsAuthors = isInAuthors(a, ResultsAuthors)
		_bookAuthor := bookAuthor.NewBook_Author(book.ISBN, _author.ID)
		ResultsBookAuth = append(ResultsBookAuth, *_bookAuthor)
	}
	return ResultsBookAuth, ResultsAuthors
}

//isInAuthors checks is in allread authors if it is not adds to authors
func isInAuthors(authr string, ResultsAuthors []author.Author) (author.Author, []author.Author) {
	for _, v := range ResultsAuthors {
		if v.NameSurname == authr {
			return v, ResultsAuthors
		}
	}
	_author := author.NewAuthor(authr)
	ResultsAuthors = append(ResultsAuthors, *_author)
	return *_author, ResultsAuthors

}
