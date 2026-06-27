package main

import (
	"encoding/json"
	"fmt"
	"os"
)


type Student struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
	Email string `json:"email"`
}

type Students struct {
	StudentList []Student `json:"studentList"`
}

//load student to reduce the duplicate LOC
func(s *Students) LoadStudent(){
		//read the existing data from the json file
	file, err := os.ReadFile("students.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	err = json.Unmarshal(file, &s)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
}

func (s *Students) SaveStudent(){
	//update the json file with the new data
	latestData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	err = os.WriteFile("students.json", latestData, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}




//add a new student to the studentList
func (s *Students) AddStudent(student Student){
	s.LoadStudent()

	s.StudentList = append(s.StudentList, student)

	s.SaveStudent()

	fmt.Println("Student added successfully!")
}

//get all students from the studentList
func (s *Students) GetAllStudents() []Student {
s.LoadStudent()

	return s.StudentList
}

//delete a student from the studentList by id
func (s *Students) DeleteStudent(id int) {
	s.LoadStudent()

	for i, student := range s.StudentList {
		if student.Id == id {
			s.StudentList = append(s.StudentList[:i], s.StudentList[i+1:]...)
			break
		}
	}

s.SaveStudent()

	fmt.Println("Student deleted successfully!")
}

//update a student in the studentList by id
func (s *Students) UpdateStudent(id int, updatedStudent Student) {
s.LoadStudent()

for i, student := range s.StudentList {
		if student.Id == id {
			s.StudentList[i] = updatedStudent
			break
		}
	}

s.SaveStudent()

	fmt.Println("Student updated successfully!")
}