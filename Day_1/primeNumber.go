package main

import "fmt"

func main() {

	var startNumber, endNumber int
	fmt.Println("Enter the start and end numbers to find prime numbers:")
	fmt.Scanln(&startNumber)
	fmt.Scanln(&endNumber)

	fmt.Print("output: ")

	for i := startNumber; i <= endNumber; i++ {
		if isPrime(i) {
			fmt.Print(i, " ")
		}
	}
}

func isPrime(number int) bool {
	for i := 2; i*i <= number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}
