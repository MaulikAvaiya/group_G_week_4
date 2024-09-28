package main

import (
	"fmt"
)

func grade() {

	subject1 := 67
	subject2 := 98
	subject3 := 77
	fmt.Println("Student Result")

	average := (subject1 + subject2 + subject3) / 3
	fmt.Println("Subject 1: ", subject1)
	fmt.Println("Subject 2: ", subject2)
	fmt.Println("Subject 3: ", subject3)
	fmt.Println("Average :", average)

	if average >= 90 {
		fmt.Println("\nExcellent! , You got Grade A")
	} else {
		if average >= 75 {
			fmt.Println("\nGood job! You got Grade B")
		} else {
			if average >= 50 {
				fmt.Println("\nYou passed. Your got Grade C")
			} else {
				fmt.Println("You are Failed and Your grade Is F")
			}
		}
	}
}
