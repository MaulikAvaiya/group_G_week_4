package main

import "fmt"

func isPrime(n int) bool {
    if n <= 1 {
        return false
    }
    for i := 2; i*i <= n; i++ {
        if n%i == 0 {
            return false
        }
    }
    return true
}

func CheckPrime() {
	fmt.Println("Enter a number to check no is prime or not: ")
	var number int
    fmt.Scan(&number)
    if isPrime(number) {
        fmt.Println(number, "is a prime number.\n")
    } else {
        fmt.Println(number, "is not a prime number.\n")
    }
}
