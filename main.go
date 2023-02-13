package main

import (
	"GoPackages/FirstWeek"
	"GoPackages/ThirdWeek"
	"fmt"
)

func main() {
	stack := &ThirdWeek.StudentStack{}

	stack.Push(&FirstWeek.Student{ID: 1, Name: "Gulnaz", Age: 18, Gender: "F", GPA: 3.7})
	stack.Push(&FirstWeek.Student{ID: 2, Name: "Nurassyl", Age: 20, Gender: "M", GPA: 3.9})
	stack.Push(&FirstWeek.Student{ID: 3, Name: "Aibek", Age: 19, Gender: "M", GPA: 3.5})

	fmt.Println("Printing stack from top to bottom:")
	stack.Print()

	fmt.Println("Printing stack from bottom to top:")
	stack.PrintReverse()

	fmt.Println("Top student:")
	stack.Peek().StudentInfo()
	fmt.Println("Popped student:")
	stack.Pop().StudentInfo()

	stack.Increment()

	fmt.Println()

	fmt.Println("Stack after incrementing ages:")
	stack.Print()

	st := &FirstWeek.Student{ID: 2, Name: "Bob", Age: 20, Gender: "M", GPA: 3.9}
	fmt.Println("Stack contains Bob:", stack.Contains(st))

	stack.Clear()

	fmt.Println("Empty stack:")
	stack.Print()
}
