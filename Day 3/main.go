package main

import (
	"fmt"
)

func main() {
	students := Students{}
	fmt.Println("Welcome to the Student Management System!")
	fmt.Println("1. Add a new Student")
	fmt.Println("2. View All students")
	fmt.Println("3. Update a student by id")
	fmt.Println("4. Delete a student by id")
	var choice int
	fmt.Print("Enter your choice: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		var newStudent Student
		fmt.Print("Enter student id: ")
		fmt.Scan(&newStudent.Id)
		fmt.Print("Enter student name: ")
		fmt.Scan(&newStudent.Name)
		fmt.Print("Enter student age: ")
		fmt.Scan(&newStudent.Age)
		fmt.Print("Enter student email: ")
		fmt.Scan(&newStudent.Email)
		students.AddStudent(newStudent)
	case 2:
		allStudents := students.GetAllStudents()
		if allStudents != nil {
			fmt.Println("List of Students:")
			for _, student := range allStudents {
				fmt.Printf("ID: %d, Name: %s, Age: %d, Email: %s\n", student.Id, student.Name, student.Age, student.Email)
			}
		}
	case 3:
		var id int
		fmt.Print("Enter student id to update: ")
		fmt.Scan(&id)
		var updatedStudent Student
		fmt.Print("Enter new Student id")
		fmt.Scan(&updatedStudent.Id)
		fmt.Print("Enter new student name: ")
		fmt.Scan(&updatedStudent.Name)
		fmt.Print("Enter new student age: ")
		fmt.Scan(&updatedStudent.Age)
		fmt.Print("Enter new student email: ")
		fmt.Scan(&updatedStudent.Email)
		students.UpdateStudent(id, updatedStudent)
	case 4:
		var id int
		fmt.Print("Enter student id to delete: ")
		fmt.Scan(&id)
		students.DeleteStudent(id)
	default:
		fmt.Println("Invalid choice. Please try again.")
		
	}
}