package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-yusufbu1ut/internal/helper/csvToDB"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-yusufbu1ut/internal/models/author"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-yusufbu1ut/internal/models/book"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-yusufbu1ut/internal/models/bookAuthor"
)

var (
	bookRepository       *book.BookRepository
	authorRepository     *author.AuthorRepository
	bookAuthorRepository *bookAuthor.BookAuthRepository
	ListErr              = errors.New("Expected List arg > 'list', 'list' 'a' or 'list' 'b'")
	SearchErr            = errors.New("Expeted Search arg > 'search' 'some arg/args for search'")
	BuyErr               = errors.New("Expected Buy args > 'buy' 'uint' 'uint'")
	buyBookErr           = errors.New("The Book is not in books")
	DeleteErr            = errors.New("Expected Delete arg> 'delete' 'uint'")
)

func init() {
	bookRepository, authorRepository, bookAuthorRepository = csvToDB.ToConnectDB() // Connecting DB and Adding sample csv datas
	//bookAuthorRepository, authorRepository, bookRepository = readInsert.ReadConnectDB() //Connecting DB, seperating CSV to slices and  Adding sample slices datas

}

func main() {

	if len(os.Args) < 2 {
		fmt.Println(ListErr)
		fmt.Println(SearchErr)
		fmt.Println(DeleteErr)
		fmt.Println(BuyErr)
		return
	}
	switch os.Args[1] {
	//List books or authors to control changings (for Buy and Delete)
	case "list":
		if len(os.Args) == 2 {
			bookList()
		} else {
			if os.Args[2] == "b" {
				if len(os.Args) == 3 {
					bookList()
				} else {
					fmt.Println(ListErr)
					return
				}
			} else if os.Args[2] == "a" {
				if len(os.Args) == 3 {
					authList()
				} else {
					fmt.Println(ListErr)
					return
				}
			} else {
				fmt.Println(ListErr)
				return
			}
		}

	//Search sends given input to funcs package Search func
	case "search":
		srch := strings.Join(os.Args[2:], " ") //Search

		if len(os.Args) > 2 {
			searchByInput(srch)
		} else {
			fmt.Println(SearchErr)
			return
		}

	// Buy comment args converting str to int and these infos goes in models package Buy func
	case "buy":
		if len(os.Args) == 4 {
			byId := os.Args[2]
			byCnt := os.Args[3]
			intId, err1 := strconv.Atoi(byId)
			intCnt, err2 := strconv.Atoi(byCnt)

			if err1 != nil || err2 != nil {
				fmt.Println(err1, err2)
				return
			}
			if intId <= 0 || intCnt <= 0 {
				fmt.Println(BuyErr)
				return
			}
			buyWithID(intId, intCnt)

		} else {
			fmt.Println(BuyErr)
			return
		}

	//Delete arg changes str to int after the process calls with Delete func for book
	case "delete":
		if len(os.Args) == 3 {
			intId, err := strconv.Atoi(os.Args[2])
			if err != nil {
				fmt.Println(err)
				return
			}
			if intId <= 0 {
				fmt.Println(DeleteErr)
				return
			}
			deleteByID(intId)

		} else {
			fmt.Println(DeleteErr)
			return
		}
	default:
		fmt.Println(ListErr)
		fmt.Println(SearchErr)
		fmt.Println(DeleteErr)
		fmt.Println(BuyErr)
		return
	}
}

// bookList Lists Books
func bookList() {
	books, err := bookRepository.GetAllBooks()

	if err != nil {
		fmt.Println(err, err.Error())
	} else {
		for _, b := range books {
			fmt.Println(b.ToString())
		}
	}
}

// authList Lists Authors
func authList() {
	authors, err := authorRepository.GetAllAuthors()

	if err != nil {
		fmt.Println(err, err.Error())
		return
	} else {
		for _, a := range authors {
			fmt.Println(a.ToString())
		}
	}
}

// searchByInput takes input parameter and first checks books and books' authors if it cant find any searchs on authors
func searchByInput(srch string) {
	books, err := bookRepository.FindByName(srch)
	//In Books
	if len(books) > 0 {
		if err != nil {
			fmt.Println(err, err.Error())
			return
		}

		for _, b := range books {
			fmt.Println("--------------------------------------------------")
			fmt.Println("Book:", b.ToString())
			book_authors, err := bookAuthorRepository.FindByISBN(b.ISBN)
			if err != nil {
				fmt.Println(err, err.Error())
				return
			} else {
				for _, ba := range book_authors {
					authors, err := authorRepository.FindByID(int(ba.AuthorID))
					if err != nil {
						fmt.Println(err, err.Error())
						return
					} else {
						for _, a := range authors {
							fmt.Println(a.ToString())
						}
					}
				}
			}
		}
	} else { // In Authors
		aurhors, err := authorRepository.FindByAuthorName(srch)
		if err != nil {
			fmt.Println(err, err.Error())
			return
		}

		for _, a := range aurhors {
			fmt.Println("--------------------------------------------------")
			fmt.Println("Author:", a.ToString())
			book_authors, err := bookAuthorRepository.FindByAuthorID(int(a.ID))
			if err != nil {
				fmt.Println(err, err.Error())
				return
			} else {
				for _, ba := range book_authors {
					books, err := bookRepository.FindByISBN(ba.BookID)
					if err != nil {
						fmt.Println(err, err.Error())
						return
					} else {
						for _, b := range books {
							fmt.Println(b.ToString())
						}
					}
				}
			}
		}
	}

}

//buyWithID works on books and takes id removes on its amount and saves it
func buyWithID(id int, cnt int) {
	book, err := bookRepository.FindByID(id)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		if len(book) == 0 {
			fmt.Println(buyBookErr)
			return
		}
		err = bookRepository.Buy(book[0], cnt)
		if err != nil {
			fmt.Println(err, err.Error())
			return
		}
	}
}

// deleteByID fun takes int count and deletes which as connected bases
func deleteByID(id int) {
	book, err := bookRepository.DeleteByID(id)

	if err != nil {
		fmt.Println(err, err.Error())
		return
	}
	authors, err := bookAuthorRepository.DeleteByISBN(book[0].ISBN)
	if err != nil {
		fmt.Println(err, err.Error())
		return
	}
	//The authors that no have in books are deleting too if he/she has no book, if delete this part only books will be deleted
	for _, a := range authors {
		err = authorRepository.DeleteByID(a)
		if err != nil {
			fmt.Println(err, err.Error())
			return
		}
	}
}
