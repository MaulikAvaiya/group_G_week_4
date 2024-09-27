package main

import "fmt"

func printFibonacci() {
	n := 10
	fmt.Println("Fibonacci Series (first", n, "terms):")
	fmt.Println("Enter the number: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Please enter a positive integer.")
		return
	}

	t1, t2 := 0, 1
	fmt.Println("Your Fibonacci Series  are: ")

	for i := 1; i <= n; i++ {
		fmt.Print(t1, " ")
		nextNumber := t1 + t2
		t1 = t2
		t2 = nextNumber
	}
	fmt.Println()
}
