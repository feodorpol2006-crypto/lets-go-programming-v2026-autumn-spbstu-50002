package main

import (
	"fmt"
	"strconv"
)

func main() {
	var firstStr, secondStr, operation string
	if _, err := fmt.Scan(&firstStr); err != nil {
		fmt.Println("Invalid first operand")
		return
	}
	first, err := strconv.Atoi(firstStr)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}
	if _, err := fmt.Scan(&secondStr); err != nil {
		fmt.Println("Invalid second operand")
		return
	}
	second, err := strconv.Atoi(secondStr)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}
	if _, err := fmt.Scan(&operation); err != nil {
		fmt.Println("Invalid operation")
		return
	}
	var result int
	switch operation {
	case "+":
		result = first + second
	case "-":
		result = first - second
	case "*":
		result = first * second
	case "/":
		if second == 0 {
			fmt.Println("Division by zero")
			return
		}
		result = first / second
	default:
		fmt.Println("Invalid operation")
		return
	}
	fmt.Println(result)
}
