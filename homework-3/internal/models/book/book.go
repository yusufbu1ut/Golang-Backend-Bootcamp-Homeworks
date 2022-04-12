package book

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-yusufbu1ut/pkg/funcs"
)

type Book struct {
	gorm.Model  //`gorm:"foreignKey:BookID;references:ID"`
	Name        string
	Pages       string
	Publisher   string
	StockCode   int     //random
	StockAmount int     //ramdom
	ISBN        int     //`gorm:"foreingKey:BookID;references:ID`
	Price       float64 //random

}

func NewBook(name string, isbn int, page string, publish string) *Book {
	return &Book{
		Name:        name,
		ISBN:        isbn,
		Pages:       page,
		Publisher:   publish,
		StockCode:   funcs.RandomInt(100000, 1000000),
		Price:       funcs.RandFloat(50, 250),
		StockAmount: funcs.RandomInt(50, 100),
	}
}

//ToString for books
func (b *Book) ToString() string {
	return fmt.Sprintf("ID : %d, Name : %s,  ISBN: %d, Stock Amount : %d, Publisher : %s", b.ID, b.Name, b.ISBN, b.StockAmount, b.Publisher)
}

//BeforeDelete sends info when deleting from db
func (b *Book) BeforeDelete(tx *gorm.DB) (err error) {
	fmt.Printf("Book (%s) deleting...", b.Name)
	return nil
}
