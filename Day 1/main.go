package main

import (
	"fmt"
)

func main(){
 for {
	var firstNumber int 
	var secondNumber int
	var calculationType string

	fmt.Println("Simple Calculator")
	fmt.Println("1.Addition \n2.Subtraction \n3.Multiplication \n4.Division")
	fmt.Scanln(&calculationType)
	fmt.Println("Enter First Number: ")
	fmt.Scanln(&firstNumber)
	fmt.Println("Enter Second Number: ")
	fmt.Scanln(&secondNumber)
	fmt.Println("You have selected:", calculationType)
	fmt.Println("The answer is:", calculate(calculationType, firstNumber, secondNumber))
	fmt.Println("Do you want to Perform another calculation? (y/n): ")
	var choice string
	fmt.Scanln(&choice)
	if choice != "y" && choice != "Y" {
		fmt.Println("Exiting the calculator. Goodbye!")
		break
	}
 }
}

func add(firstNumber int, secondNumber int) int {
	return firstNumber + secondNumber
}
func subtract(firstNumber int, secondNumber int) int {
	return firstNumber - secondNumber
}
func multiply(firstNumber int, secondNumber int) int {
	return firstNumber * secondNumber
}
func divide(firstNumber int, secondNumber int) float64 {
	if secondNumber == 0 {
		fmt.Println("Error: Division by zero is not allowed.")
		return 0
	}
	return float64(firstNumber) / float64(secondNumber)
}

func calculate(calculationType string, firstNumber int, secondNumber int) interface{} {
	switch calculationType {
	case "1":
		return add(firstNumber, secondNumber)
	case "2":
		return subtract(firstNumber, secondNumber)
	case "3":
		return multiply(firstNumber, secondNumber)
	case "4":
		return divide(firstNumber, secondNumber)
	default:
		return "Invalid calculation type selected."
	}
}
