package main

import (
	"bufio" //dosya okurken kullanılıyor
	"fmt"
	"os"
	"strings"

	"github.com/yusufbu1ut/KitaplikApp/helper"
)

var BookSlice []string

func init() {
	//Dosya okunmak üzere açılır
	//dizindeki dosyanın bulunamaması durumunda panic hata çıktısı verilir
	file, err := os.OpenFile("kitaplar.txt", os.O_RDONLY, 0755)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	defer file.Close()

	for scanner.Scan() {
		//burada gerçekleştirilen okuma değerleri slice içerisine aktarılır
		//kitaplar slice içerisine yerleştirilir burada girdi sayısının belirsiz olduğu var sayılmıştır
		line := scanner.Text()
		BookSlice = append(BookSlice, line)
	}
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Expected 'search' or 'list'")
	}
	//fmt.Println(os.Args)
	switch os.Args[1] {
	case "list":
		if len(os.Args) >= 3 && os.Args[2] != "" {
			fmt.Println("Command 'list' doesnt take any arg")
		} else {
			helper.List(BookSlice)
		}
	case "search":
		srch := strings.Join(os.Args[2:], " ") //Search argumanları birleştirilir
		if len(os.Args) > 2 {
			helper.Search(srch, BookSlice)
		} else {
			fmt.Println("Expected search argument for command 'search'")
		}
	default:
		fmt.Println("Expected 'list' or 'serach'")
	}

	println("")
}
