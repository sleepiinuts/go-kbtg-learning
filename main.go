package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	models "github.com/sleepiinuts/go-kbtg-learning/models"
)

type Address struct {
	HomeNo     string
	StreetName string
	Province   string
	Zipcode    int
}

func (a *Address) String() string {
	aValue := reflect.ValueOf(*a)
	aType := aValue.Type()
	i := aValue.NumField()

	str := "{\n"

	for idx := range i {
		str += fmt.Sprintf("   %s: %v\n", aType.Field(idx).Name, aValue.Field(idx))
	}

	return str + "}"
}

func main() {
	str := "golang is a simple programming language"
	i := 1024
	f64 := float64(100.35)
	b := true
	bb := []byte(`simple byte string`)

	fmt.Printf("%-20v -- type: %T\n", str, str)
	fmt.Printf("%-20v -- type: %T\n", i, i)
	fmt.Printf("%-20v -- type: %T\n", f64, f64)
	fmt.Printf("%-20v -- type: %T\n", b, b)
	fmt.Printf("%-20v -- type: %T\n", string(bb), bb)

	j := 10
	jj := float64(j)
	fmt.Printf("%v:%v\n", &j, &jj)

	addr := Address{
		HomeNo:     "home",
		StreetName: "street",
		Province:   "province",
		Zipcode:    112,
	}

	addr2 := Address{
		// HomeNo:     "home2",
		StreetName: "street2",
		Province:   "province2",
		Zipcode:    112,
	}

	// addr2 = addr
	fmt.Println("reassign: ", &addr)
	fmt.Println("reassign: ", &addr2)

	// fmt.Printf("addr: %v\n", addr)
	if jstr, err := json.MarshalIndent(addr, "", "    "); err == nil {
		fmt.Println(string(jstr))
	}

	// fmt.Printf("%+v", addr)

	// fmt.Println(addr)
	fmt.Println("stringer: ", &addr)

	citizenID := []int{1, 1, 0, 3, 0, 0, 0, 0, 1, 9, 7, 8, 0}
	fmt.Println("my last 4 digits citizenID: ", citizenID[len(citizenID)-4:])

	cID := strings.Split(`1103000019780`, "")
	fmt.Println("my last 4 digits cID: ", cID[len(cID)-4:])

	printScore(100)

	printEvery5()

	fmt.Println("----")

	printEven()

	// defer func() {
	// 	str := recover()
	// 	fmt.Println("recover: ", str)
	// 	debug.PrintStack()
	// }()

	// panic("some message")

	c := models.Car{
		Name:  "sample name",
		Model: "sample model",
		Price: 123.45,
	}

	models.PrintDetails(c.Name, c.Model, c.Price)
	c.PrintDetaisWithCustomType()
	c.PrintDetailsWithCustomTypePretty()

	fmt.Println(models.PrintThousands(12345))

	e := models.Employee{
		FirstName: "FName",
		LastName:  "LName",
		Salary:    1234.56,
	}

	e.Print()
	e.PrintSalary()

	fmt.Println(sumComparable([]int{1, 2, 3, 4, 5}))
	fmt.Println(sumComparable([]float32{1.1, 1.2, 2.3}))
}

func printEvery5() {
	for i := 0; i <= 100; i += 5 {
		fmt.Printf("i: %d\n", i)
	}
}

func printEven() {
	for i := range 101 {
		if i%2 == 0 {
			fmt.Printf("i: %d\n", i)
		}
	}
}

func printScore(scr int) {
	switch {
	case scr > 90:
		fmt.Println("more than 90")
	case scr > 80:
		fmt.Println("more than 80")
	case scr > 70:
		fmt.Println("more than 70")
	case scr > 60:
		fmt.Println("more than 60")
	default:
		fmt.Println("else")
	}
}

func init() {
	fmt.Println("init function")
}
