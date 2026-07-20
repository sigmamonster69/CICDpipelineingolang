package main

import "fmt"

// Add takes two integers and returns their sum
func Add(a, b int) int {
	return a + b
}

// Subtract takes two integers and returns their difference
func Subtract(a, b int) int {
	return a - b
}

func main() {
	fmt.Println("Welcome to our CI/CD Demo App!")
	fmt.Printf("5 + 3 = %d\n", Add(5, 3))
	fmt.Printf("10 - 4 = %d\n", Subtract(10, 4))
}
