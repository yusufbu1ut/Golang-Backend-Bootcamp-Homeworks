package helper

import (
	"fmt"
	"strings"
)

func Search(input string, list []string) {
	fmt.Printf("\n'%s' Içeren Kitaplar\n----------------------------------------\n", input)
	for i, v := range list {
		//kitaplar içerisinde girilen değeri içeren kitaplar ve slice içerisindeki konumu çıktı olarak yazdırılır
		//karşılaştırma girilen değer ve kitap adaları küçük,büyük harfe döndürülerek  karşılaştırmalar gerçekleştirilir
		if strings.Contains(strings.ToLower(v), strings.ToLower(input)) || strings.Contains(strings.ToUpper(v), strings.ToUpper(input)) {
			fmt.Printf("%d %s\n", i+1, v)
		}
	}
}

func List(list []string) {
	//parametresi slice olan girdi içerisinde gezilerek çıktı verilir
	fmt.Printf("\nKitaplık\n----------------------------------------\n")
	for i, v := range list {
		fmt.Printf("%d %s\n", i+1, v)
	}
}
