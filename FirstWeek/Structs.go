package FirstWeek

import "fmt"

type Student struct {
	ID      int
	Name    string
	Age     int
	Gender  string
	GPA     float64
	Subject []Subject
	Grade   []Grade
}

type School struct {
	ID       int
	Name     string
	Location string
	Students []Student
}

type Subject struct {
	ID      int
	Name    string
	Teacher string
}

type Grade struct {
	StudentID int
	SubjectID int
	Grade     float64
}

func (s *Student) StudentInfo() {
	fmt.Println("Student name:", s.Name)
	fmt.Println("Student id", s.ID)
	fmt.Println("Student age", s.Age)
	printGrade(s)
}

func printGrade(s *Student) {
	gpa := 0.0
	for i, grade := range s.Grade {

		fmt.Println(s.Subject[i].Name, grade.Grade)
		gpa = gpa + grade.Grade
	}
	fmt.Println("Student's average GPA is: ", gpa/float64(len(s.Subject)))
}
