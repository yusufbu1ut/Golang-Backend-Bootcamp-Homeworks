package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-2-yusufbu1ut/helper"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-2-yusufbu1ut/models"
)

var slcBook []models.Book

func init() {
	//Adding book informations to slice
	a1 := models.NewAuthor("George", "Orwell")
	b1 := models.NewBook("1984", *a1)
	slcBook = append(slcBook, *b1)

	a2 := models.NewAuthor("Alfredo", "Covelli")
	b2 := models.NewBook("Vahana Masterclass", *a2)
	slcBook = append(slcBook, *b2)

	a3 := models.NewAuthor("Hunter", "Biden")
	b3 := models.NewBook("Beautiful Things’ A Memoir", *a3)
	slcBook = append(slcBook, *b3)

	a4 := models.NewAuthor("Dalai", "Lama")
	b4 := models.NewBook("The Little Book of Encouragement", *a4)
	slcBook = append(slcBook, *b4)

	a5 := models.NewAuthor("Ramachandra", "Guha")
	b5 := models.NewBook("The Commonwealth of Cricket", *a5)
	slcBook = append(slcBook, *b5)

	b6 := models.NewBook("Diaries", *a1)
	slcBook = append(slcBook, *b6)

	a6 := models.NewAuthor("Platon", "")
	b7 := models.NewBook("Parmenides", *a6)
	slcBook = append(slcBook, *b7)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println(models.ExpectedDef.Error())
	}
	switch os.Args[1] {
	//Listing books to control changings (for Buy and Delete)
	case "list":
		if len(os.Args) == 2 {
			helper.List(slcBook)
		} else {
			fmt.Println(models.ErrList.Error())
		}
	// In buy comment args converting str to int and these infos goes in models package Buy func
	case "buy":
		if len(os.Args) == 4 {
			byId := os.Args[2]
			byCnt := os.Args[3]
			intId, err1 := strconv.Atoi(byId)
			intCnt, err2 := strconv.Atoi(byCnt)

			if err1 != nil || err2 != nil {
				fmt.Println(models.ExpectedBuyArg.Error())
			}
			for i, book := range slcBook {
				if book.Id == intId {
					slcBook[i].Buy(intCnt)
					break
				}
				if i == len(slcBook)-1 && book.Id != intId {
					fmt.Println(models.ExpectedValidBuy.Error())
				}
			}

		} else {
			fmt.Println(models.ExpectedBuy.Error())
		}
	//Search sends given input to funcs package Search func
	case "search":
		srch := strings.Join(os.Args[2:], " ") //Search
		if len(os.Args) > 2 {
			helper.Search(srch, slcBook)
		} else {
			fmt.Println(models.ExpectedSearchArg.Error())
		}
	//Delete arg changes str to int after the process calls with Delete func for book
	case "delete":
		if len(os.Args) == 3 {
			byId := os.Args[2]
			intId, err1 := strconv.Atoi(byId)
			if err1 != nil {
				fmt.Println(models.ExpectedDeleteArg.Error())
			}
			for i, b := range slcBook {
				if b.Id == intId {
					slcBook[i].Delete()
					break
				}
				if i == len(slcBook)-1 && b.Id != intId {
					fmt.Println(models.ExpectedValidDelete.Error())
				}

			}

		} else {
			fmt.Println(models.ExpectedDelete.Error())
		}
	default:
		fmt.Println(models.ExpectedDef.Error())
	}

	println("")
}
