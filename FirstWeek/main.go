package FirstWeek

import (
	"fmt"
	"math/rand"
)

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

	student1.StudentInfo()
	fmt.Println("----------------------")
	student2.StudentInfo()

}
