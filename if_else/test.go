package main

import "fmt"

func main() {
	var name string
	print("Enter your name account: ")
	fmt.Scanf("%s", &name) // %s is a format specifier for string
	fmt.Println("Hello", name)
	
	var score int 
	fmt.Println("Enter your score: ")
	fmt.Scanf("%d", &score) // %d is a format specifier for int
	if score <= 50 {
		fmt.Println("You need to improve your score.")
	} else if score > 50 && score <= 75 {
		fmt.Println("You are doing well, keep it up!")
	} else {
		fmt.Println("Excellent work, you are a star!")
	}
}