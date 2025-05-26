package main

import "fmt"

// add parameters unlimit

func main() {
	result := variadic(10, 20, 30, 40, 50)
	fmt.Println("The result of the number function is", result)
}

func variadic(numbers ...int) int { // numbers is a variadic parameter
	total := 0
	for _, value := range numbers {
		total += value // sum all numbers
	}
	return total // return the total
}