package main

import "fmt"

func main() {
	numbers := []int{100,200,300}
	numbers = append(numbers, 400)
	fmt.Println("The numbers are", numbers)
	
}