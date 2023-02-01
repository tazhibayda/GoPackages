package main

import (
	"fmt"
	"math/rand"
)

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

func main() {
	golang := Subject{Name: "The Go programming language", Teacher: "Azamat Serek"}
	javaee := Subject{Name: "Advanced Programming using JAVA", Teacher: "Yerlan Karabaliyev"}
	stat := Subject{Name: "Probability and Mathematical Statictics", Teacher: "Larissa Bazarbayeva"}

	student1 := Student{ID: 200103210, Name: "Daryn Tazhibay", Age: 19, Gender: "Male", GPA: 0, Subject: []Subject{golang, javaee, stat}}
	student2 := Student{ID: 200000001, Name: "Aruzhan Aaa", Age: 19, Gender: "Female", GPA: 0, Subject: []Subject{golang, javaee, stat}}

	school := School{Name: "SDU", Location: "Qaskelen", Students: []Student{student1, student2}}

	student1.Grade = make([]Grade, len(student1.Subject))
	student2.Grade = make([]Grade, len(student2.Subject))

	for i, subject := range student1.Subject {

		student1.Grade[i] = Grade{
			StudentID: student1.ID,
			SubjectID: subject.ID,
			Grade:     (float64(rand.Intn(300.0)) + 100) / 100,
		}

	}
	for i, subject := range student2.Subject {

		student2.Grade[i] = Grade{
			StudentID: student1.ID,
			SubjectID: subject.ID,
			Grade:     (float64(rand.Intn(300.0)) + 100) / 100,
		}

	}

	fmt.Println("School Name:", school.Name)
	fmt.Println("Location:", school.Location)
	fmt.Println("Number of Students:", len(school.Students))

	student1.studentInfo()
	fmt.Println("----------------------")
	student2.studentInfo()

}

func (s *Student) studentInfo() {
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
