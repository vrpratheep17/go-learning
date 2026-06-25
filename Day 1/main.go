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
	result, err := calculate(calculationType, firstNumber, secondNumber)
	if err != nil {
	fmt.Println(err)
	} else {
	fmt.Println("Result:", result)
	}
	fmt.Println("Do you want to Perform another calculation? (y/n): ")
	var choice string
	fmt.Scanln(&choice)
	if choice != "y" && choice != "Y" {
		fmt.Println("Exiting the calculator. Goodbye!")
		break
	}
 }
}

func add(firstNumber int, secondNumber int) float64 {
	return float64(firstNumber) + float64(secondNumber)
}
func subtract(firstNumber int, secondNumber int) float64 {
	return float64(firstNumber) - float64(secondNumber)
}
func multiply(firstNumber int, secondNumber int) float64 {
	return float64(firstNumber) * float64(secondNumber)
}
func divide(firstNumber int, secondNumber int) float64 {
	if secondNumber == 0 {
		fmt.Println("Error: Division by zero is not allowed.")
		return 0
	}
	return float64(firstNumber) / float64(secondNumber)
}

func calculate(calculationType string, firstNumber int, secondNumber int) (float64,error) {
	switch calculationType {
	case "1":
		return add(firstNumber, secondNumber), nil
	case "2":
		return subtract(firstNumber, secondNumber), nil
	case "3":
		return multiply(firstNumber, secondNumber), nil
	case "4":
		return divide(firstNumber, secondNumber), nil
	default:
		return 0, fmt.Errorf("invalid calculation type")
	}
}
