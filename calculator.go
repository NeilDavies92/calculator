package main

import (
	"fmt"
)

func main() {

	var num1, num2 float64
	var operator string

	fmt.Print("Please enter a number: ")
	fmt.Scanln(&num1)

	fmt.Print("Please enter a second number: ")
	fmt.Scanln(&num2)

	fmt.Print("Please enter an operator: ")
	fmt.Scanln(&operator)

	switch operator {
	case "+":
		fmt.Printf("%f %s %f = %f", num1, operator, num2, num1+num2)
	case "-":
		fmt.Printf("%f %s %f = %f", num1, operator, num2, num1-num2)
	case "*":
		fmt.Printf("%f %s %f = %f", num1, operator, num2, num1*num2)
	case "/":
		if num2 == 0 {
			fmt.Print("Divide by 0 is not possible")
		} else {
			fmt.Printf("%f %s %f = %f", num1, operator, num2, num1/num2)
		}
	default:
		fmt.Println("Invalid Operator")
	}
}
