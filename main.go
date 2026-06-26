package main

import "fmt"

func main() {

	fmt.Println("Welcome to Go Calculator.")

	for {

		fmt.Print("Enter the operator from ( + ,- ,* , / ): ")
		var operator string
		fmt.Scan(&operator) // operator input for operations

		if operator == "q" {
			fmt.Println("Good Bye.!!") // ending functionality
			break
		}

		fmt.Print("Enter the Number 1: ") // First operant
		var num1 float64
		fmt.Scan(&num1)

		fmt.Print("Enter the Number 2: ") // Second operand
		var num2 float64
		fmt.Scan(&num2)

		result, err := calculate(num1, num2, operator)

		if err == nil {
			fmt.Printf("The result for %v %s %v = %v \n", num1, operator, num2, result)
		} else {
			fmt.Printf("Error: %v .\n", err)
		}
	}
}

func calculate(a float64, b float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("cannot divide by zero")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("unrecognized operator: %s", operator)
	}
}
