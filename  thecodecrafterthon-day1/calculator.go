package main

import (
	"fmt"
)

func main() {
	for {
		var input1 int
		var input2 int
		var operator string
		var choice string

		fmt.Print("Enter user input: ")
		fmt.Scan(&input1)

		fmt.Print("Enter user input: ")
		fmt.Scan(&input2)

		fmt.Println("<1> addition")
		fmt.Println("<2> subtraction")
		fmt.Println("<3> multiplication")
		fmt.Println("<4> division")
		fmt.Println("<5> quit")

		fmt.Print("Select Operator: ")
		fmt.Scan(&operator)

		if operator == "quit" {
			fmt.Print("Good bye")
			break
		}
		if operator == "help" {
			fmt.Println("Unknown command. Expected command")
			fmt.Println("<1> add,<2> sub,<3> mul, <4> div,")
		}
		switch operator {
		case "1":
			fmt.Println("Result =", input1+input2)

		case "2":
			fmt.Println("Result =", input1-input2)

		case "3":
			fmt.Println("Result =", input1*input2)

		case "4":
			if input2 == 0 {
				fmt.Println("Error: Division by zero is not allowed")
			} else {
				fmt.Println("Result =", input1/input2)
			}

		default:
			fmt.Println("Invalid operator selected")
			fmt.Scan(&choice)
			continue
		}

		fmt.Println("Enter user input: ")
		fmt.Scan(&choice)

		if choice == "n" || choice == "N" {
			fmt.Println("Calculator exited.")
			break
		}
	}
}
