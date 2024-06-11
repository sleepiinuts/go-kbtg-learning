package models

import (
	"fmt"

	"github.com/leekchan/accounting"
)

type Employee struct {
	FirstName string
	LastName  string
	Salary    float32
}

func (e *Employee) PrintSalary() {
	ac := accounting.Accounting{Symbol: "$", Precision: 3, Thousand: ","}
	fmt.Printf("%s %s's salary is %s\n", e.FirstName, e.LastName, ac.FormatMoney(e.Salary))
}

func (e *Employee) Print() {
	fmt.Printf("%+v\n", *e)
}
