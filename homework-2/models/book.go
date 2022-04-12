package models

import (
	"fmt"
)

var id int = 0                //for id increase +1 when added new item Book
var stock_code int = 10000000 //for stock_code increase +1 when added new item Book
//Book contains Author infos
type Book struct {
	IsDeleted                             bool
	Id, StockCode, ISBN, StockNo, PagesNo int
	Price                                 float64
	Name                                  string
	Author
}

//Author
type Author struct {
	isDeleted bool
	//id            int
	name, surname string
}

//constructor for Author
func NewAuthor(name string, surname string) *Author {

	return &Author{
		name:      name,
		surname:   surname,
		isDeleted: false,
	}
}

//Author print func brings together author infos
func (auth *Author) Print() string {
	return auth.name + " " + auth.surname
}

//constructor for Book
func NewBook(newName string, writer Author) *Book {

	id += 1
	stock_code += 1

	return &Book{
		Name:      newName,
		PagesNo:   RandomInt(50, 500), //random
		Price:     RandFloat(15, 250), //random
		StockNo:   RandomInt(15, 20),  //random
		Author:    writer,
		ISBN:      RandomInt(1000000000, 10000000000), //random
		Id:        id,
		StockCode: stock_code,
		IsDeleted: false,
	}
}

type Deletable interface {
	Delete()
}

//delete func for Book returns changed book
func (book *Book) Delete() {
	if !book.IsDeleted {
		book.IsDeleted = true
	} else {
		fmt.Println(NotInbooks.Error())
	}
}

//Buy func comes buyying count and process on stock number returns changed book infos
func (book *Book) Buy(count int) {
	if book.StockNo < count {
		fmt.Println(HigherThanStock.Error())
	} else if book.IsDeleted {
		fmt.Println(NotInbooks.Error())
	} else {
		book.StockNo = book.StockNo - count
	}
}
