package samples

import "fmt"

/*
Ideas for code snippets
Hard to master: Allocation with make
Verbose Error Handling: Handling sql errors after multiple
*/

func main() {
	// Print a simple "Hello, Go!" message
	fmt.Println("Hello, Go!")

	// Declare and initialize variables
	age := 30
	name := "John Doe"
	fmt.Printf("Name: %s, Age: %d\n", name, age)

	// Use a for loop to print numbers from 1 to 5
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// Define a function that adds two numbers
	add := func(a, b int) int {
		return a + b
	}

	// Call the function and print the result
	result := add(3, 7)
	fmt.Println("3 + 7 =", result)
}
