package main

import (
	"fmt"
)

func main() {

	var num1, num2 float64
	var opperator string

	fmt.Print("Please Enter a number: ")
	fmt.Scanln(&num1)

	fmt.Print("Please Enter a second number: ")
	fmt.Scanln(&num2)

	fmt.Print("Please enter an opperator: ")
	fmt.Scanln(&opperator)

	switch opperator {
	case "+":
		fmt.Printf("%f %s %f = %f", num1, opperator, num2, num1+num2)
	case "-":
		fmt.Printf("%f %s %f = %f", num1, opperator, num2, num1*num2)
	case "*":
		fmt.Printf("%f %s %f = %f", num1, opperator, num2, num1+num2)
	case "/":
		if num2 == 0 {
			fmt.Print("Divide by 0 is not possible")
		} else {
			fmt.Printf("%f %s %f = %f", num1, opperator, num2, num1/num2)
		}
	}
}
