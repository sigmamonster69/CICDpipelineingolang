package main

import "fmt"

// Add takes two integers and returns their sum
// This function will be automatically tested by our CI pipeline
// If you change this logic, the tests will catch it!
func Add(a, b int) int {
	return a + b
}

// Subtract takes two integers and returns their difference
// Same here - any bugs introduced will be caught by automated tests
func Subtract(a, b int) int {
	return a - b
}

// main is the entry point of our application
// When deployed, this is what runs on the server
func main() {
	// Welcome message - you'll see this when the app runs
	fmt.Println("Welcome to our CI/CD Demo App!")
	
	// Demonstrate the Add function
	fmt.Printf("5 + 3 = %d\n", Add(5, 3))
	
	// Demonstrate the Subtract function
	fmt.Printf("10 - 4 = %d\n", Subtract(10, 4))
}
