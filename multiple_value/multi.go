package main

import "fmt"

func main() {
	result, check := summation(10, 20)
	fmt.Println("The result of the summation is", result)
	fmt.Println("Check status:", check)
}

func summation(num1, num2 int) (int, string) { // int is the return type
	total := num1+num2
	status := "" 
	if total%2 == 0 {
		status = "even"
	} else {
		status = "odd"
	}
	return total, status
}