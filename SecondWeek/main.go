package SecondWeek

import (
	"fmt"
	"github.com/tazhibayda/Golang/Companies"
	"github.com/tazhibayda/Golang/Math"
)

func main() {

	var square Math.Square
	var company Companies.Company
	var persons []Companies.Person

	square = Math.Square{AC: 4, BD: 4}

	fmt.Println(square.Info())
	fmt.Println("-----------------------------")
	fmt.Println(Math.GetArea(square.BD, square.BD))

	fmt.Println("-----------------------------")
	company = Companies.Company{
		Name:       "B7",
		Experience: 4,
		Employees:  nil,
	}

	fmt.Println(company.CompanyInfo())
	fmt.Println("-----------------------------")

	persons = make([]Companies.Person, 4)

	persons[0] = Companies.Person{
		Name:       "Aibek",
		Age:        25,
		Experience: 7,
	}
	persons[1] = Companies.Person{
		Name:       "Nurzhan",
		Age:        32,
		Experience: 2,
	}
	persons[2] = Companies.Person{
		Name:       "Aidos",
		Age:        18,
		Experience: 3,
	}
	persons[3] = Companies.Person{
		Name:       "Assyl",
		Age:        23,
		Experience: 11,
	}

	for _, person := range persons {
		fmt.Println(person.PersonInfo())
		fmt.Println(company.Hire(person))
		fmt.Println("-----------------------------")
	}

	for _, person := range company.CompanyEmployees() {
		fmt.Println(person.PersonInfo())
	}

}
