package main

import "fmt"

func Pyramid() {
	var a int
	fmt.Print("Enter the number of rows: ")
	fmt.Scan(&a)

	// for generating pyramid
	for i := 1; i <= a; i++ {
		// for printing spaces
		for j := 1; j <= a-i; j++ {
			fmt.Print(" ")
		}
		// for printing stars
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		// for moving to next line after printing each row
		fmt.Println()
	}
}
