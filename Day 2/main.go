package main

import (
	"fmt"
)

type Student struct {
	ID     int
	Name   string
	Age    int
	Course string
}

type StudentService struct {
	students []Student
}

func (s *StudentService) AddStudent(student Student) {
	s.students = append(s.students, student)
}

func (s *StudentService) GetStudents() []Student {
	return s.students
}

func (s *StudentService) DeleteStudent(id int) {
	for i, student := range s.students {
		if student.ID == id {
			s.students = append(s.students[:i], s.students[i+1:]...)
			break
		}
	}
}

func (s *StudentService) UpdateStudent(id int, updatedStudent Student){
	for i, student := range s.students {
		if student.ID == id {
			s.students[i] = updatedStudent
			break
		}
	}
}

func main() {
	studentService := StudentService{}

	student1 := Student{ID: 1, Name: "Alice", Age: 20, Course: "Math"}
	student2 := Student{ID: 2, Name: "Bob", Age: 22, Course: "Science"}
	choice := 0

	studentService.AddStudent(student1)
	studentService.AddStudent(student2)

	for {
		fmt.Println("===== Student Management =====")
		fmt.Println("1. Add Student")
		fmt.Println("2. View Students")
		fmt.Println("3. Update Student")
		fmt.Println("4. Delete Student")
		fmt.Println("5. Exit")
		fmt.Println()
		fmt.Print("Choose:")

		fmt.Scanln(&choice)

	switch choice {
	case 1:
		println("Enter Student ID:")
		var id int
		fmt.Scanln(&id)
		println("Enter Student Name:")
		var name string
		fmt.Scanln(&name)
		println("Enter Student Age:")
		var age int
		fmt.Scanln(&age)
		println("Enter Student Course:")
		var course string
		fmt.Scanln(&course)
		student := Student{ID: id, Name: name, Age: age, Course: course}
		studentService.AddStudent(student)
	case 2:
		students := studentService.GetStudents()
		for _, student := range students {
			fmt.Printf("ID: %d, Name: %s, Age: %d, Course: %s\n", student.ID, student.Name, student.Age, student.Course)
		}

	case 3:
		println("Enter Student ID to Update:")
		var id int
		fmt.Scanln(&id)
		println("Enter Updated Student Name:")
		var name string
		fmt.Scanln(&name)
		println("Enter Updated Student Age:")
		var age int
		fmt.Scanln(&age)
		println("Enter Updated Student Course:")
		var course string
		fmt.Scanln(&course)
		updatedStudent := Student{ID: id, Name: name, Age: age, Course: course}
		studentService.UpdateStudent(id, updatedStudent)

	case 4:
		println("Enter Student ID to Delete:")
		var id int
		fmt.Scanln(&id)
		studentService.DeleteStudent(id)
	case 5:
		println("Exiting...")
		return
	default:
		println("Invalid choice. Please try again.")
	}
}
}