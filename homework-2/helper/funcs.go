package helper

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-2-yusufbu1ut/models"
)

//Search func firstly convert given input str to int for sku infos
//given str first compare with book names with tolower
//secondly compare with Author.Print which as author name surname with tolower
//lastly converted str to int info compare with book stock code or number
//but this process needs full of code or piece number so it works if user enters correctly full of stock code or full of stock number
func Search(input string, slice []models.Book) {
	inp, _ := strconv.Atoi(input)
	fmt.Printf("Results for count '%v' \n", input)
	fmt.Println("-----------------------------------")
	for _, b := range slice {
		if strings.Contains(strings.ToLower(b.Name), strings.ToLower(input)) && !b.IsDeleted {
			fmt.Println(b.Id, " ", b.StockCode, " ", b.Name, " ", b.Author.Print(), " ", b.StockNo)
		} else if strings.Contains(strings.ToLower(b.Author.Print()), strings.ToLower(input)) && !b.IsDeleted {
			fmt.Println(b.Id, " ", b.StockCode, " ", b.Name, " ", b.Author.Print(), " ", b.StockNo)
		} else if (inp == b.StockCode || inp == b.StockNo) && !b.IsDeleted {
			fmt.Println(b.Id, " ", b.StockCode, " ", b.Name, " ", b.Author.Print(), " ", b.StockNo)
		}
	}
}

//List lists Book slice infos I used this to check delete and buy processes
func List(slice []models.Book) {
	fmt.Println("Books")
	fmt.Println("-----------------------------------")
	for _, b := range slice {
		if !b.IsDeleted {
			fmt.Println(b.Id, " ", b.Name, " ", b.Author.Print(), " ", b.StockNo, " ", b.IsDeleted, " ", b.Price, " ", b.ISBN)
		}
	}
}
