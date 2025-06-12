package main

import "fmt"

func main() {
	scanf1() 
	scanf2()
}

func scanf2() {
	// var score int
	// fmt.Print("Enter your score: " )
	// fmt.Scanf("%d", &score) // %d is a format specifier for int
	// fmt.Println("Your score is", score)

	// //Process the score
	// if score >= 50 {
	// 	fmt.Println("You passed")
	// } else {
	// 	fmt.Println("You failed")
	// }

	var num int
	fmt.Print("Enter a number: ")
	fmt.Scanf("%d", &num) // %d is a format specifier for int

	if num%2 == 0 {
		fmt.Println("The number is even")
	} else {
		fmt.Println("The number is odd")
	}

}

func scanf1() {
	var name string
	var score int
	fmt.Print("Enter your name: ")
	fmt.Scanf("%s", &name) // %s is a format specifier for string
	fmt.Print("Enter your score: ")
	fmt.Scanf("%d", &score) // %d is a format specifier for int
	fmt.Println("Hello", name)
	fmt.Println("Your score is", score+10)
}

