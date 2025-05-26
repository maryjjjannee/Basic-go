package main

import "fmt"

func main() {
	
	var name string
	fmt.Print("Enter your name account: ")
	fmt.Scanf("%s", &name) // %s is a format specifier for string
	fmt.Println("Hello", name)

	var todo int
	fmt.Println("Enter 1 to deposit, 2 to withdraw")
	fmt.Scanf("%d", &todo) // %d is a format specifier for int

	if todo == 1 {
		var deposit int
		fmt.Print("Enter amount to deposit: ")
		fmt.Scanf("%d", &deposit)
		fmt.Println("You deposited", deposit)
	} else if todo == 2 {
		var withdraw int
		fmt.Print("Enter amount to withdraw: ")
		fmt.Scanf("%d", &withdraw)
		fmt.Println("You withdrew", withdraw)
	}
}
