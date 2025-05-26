package main

import "fmt"

func main() {
	showMessage("Jane!")
	total(28, 5)
	delivery := getdelivery()
	fmt.Println("The delivery fee is", delivery)
	myCart := cart(100, 200)
	fmt.Println("The total amount in the cart is", myCart)
}

func cart(num1, num2 int) int {
	total := num1+num2
	return total
}

func getdelivery() int { 
	return 50
}

func total(num1 int, num2 int){
	fmt.Println("The total is", num1 + num2)
}

func showMessage(name string) {
	fmt.Println("Hello", name)
}