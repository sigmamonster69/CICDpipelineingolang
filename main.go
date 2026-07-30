package main

import "fmt"

// Add takes two integers and returns their sum.
// This function demonstrates basic arithmetic operations in Go.
// The CI pipeline automatically tests this function on every push.
// If you modify this logic, the automated tests will catch any regressions.
// Parameters:
//   - a: first integer operand
//   - b: second integer operand
// Returns:
//   - int: the sum of a and b
func Add(a, b int) int {
	return a + b
}

// Subtract takes two integers and returns their difference (a - b).
// This function is part of our demo to show how CI/CD catches bugs.
// The automated test suite validates this function runs correctly.
// Parameters:
//   - a: minuend (the number from which another number is subtracted)
//   - b: subtrahend (the number to be subtracted)
// Returns:
//   - int: the result of subtracting b from a
func Subtract(a, b int) int {
	return a - b
}

// main is the entry point of our application.
// When the application is deployed via CI/CD, this function executes first.
// It demonstrates the core functionality of our demo app by calling
// the Add and Subtract functions and printing results to stdout.
// In production, this would typically start a web server or service.
func main() {
	// Print welcome message to indicate the application has started successfully
	// This message confirms the binary was built and is running correctly
	fmt.Println("Welcome to our CI/CD Demo App!")
	
	// Demonstrate the Add function with sample values (5 + 3)
	// The output helps verify the function works as expected during manual testing
	fmt.Printf("5 + 3 = %d\n", Add(5, 3))
	
	// Demonstrate the Subtract function with sample values (10 - 4)
	// This provides visual confirmation that the subtraction logic is correct
	fmt.Printf("10 - 4 = %d\n", Subtract(10, 4))
}
