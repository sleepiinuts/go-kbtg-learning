package models

import (
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type Car struct {
	Name  string
	Model string
	Price float32
}

func PrintDetails(name, model string, price float32) {
	fmt.Println("PrintDetails")
	fmt.Println("Name: ", name)
	fmt.Println("Model: ", model)
	fmt.Println("price: ", price)
	fmt.Println("----")
}

func (c *Car) PrintDetaisWithCustomType() {
	fmt.Println("PrintDetaisWithCustomType")
	fmt.Printf("%+v\n", *c)
	fmt.Println("----")
}

func (c *Car) PrintDetailsWithCustomTypePretty() {
	fmt.Println("PrintDetailsWithCustomTypePretty")
	fmt.Println("Name: ", c.Name)
	fmt.Println("Model: ", c.Model)

	price := message.NewPrinter(language.English)
	price.Printf("Price: $%d\n", 1000)

	fmt.Println("----")
}

func PrintThousands(n int) string {
	fmt.Println("PrintThousands")
	str := ""
	for {
		str = "," + strconv.Itoa(n%1000) + str
		// fmt.Println("str: ", str)
		if n < 1000 {
			break
		}
		n /= 1000
	}
	return strings.Trim(str, ",")
}
