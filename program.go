package main

import (
	"fmt"
	"gobasic/calculator"
)

type Product struct {
	name string
	price float64
	category string
	discount int
}

func main() {
	// profile()
	// math1()
	// math2()
	// scanf1()
	// scanf2()
	// bank()
	// slice()
	// array()
	// map1()
	// loop()
	// Rage()
	// showMessage("Jane!")
	// total(28, 5)
	// delivery := getdelivery()
	// fmt.Println("The delivery fee is", delivery)
	// myCart := cart(100, 200)
	// fmt.Println("The total amount in the cart is", myCart)
	// result, check := summation(10, 20)
	// fmt.Println("The result of the summation is", result)
	// fmt.Println("Check status:", check)
	// result := number(10, 20, 30, 40, 50)
	// fmt.Println("The result of the number function is", result)
	// product1 := Product{name: "Laptop", price: 50000, category: "Electronics", discount: 10}
	// product2 := Product{name: "Mouse", price: 1000, category: "Electronics", discount: 20}
	// fmt.Println(product1)
	// fmt.Println(product2)
	fmt.Println(calculator.Add(10, 20)) // call the Add function from the calculator package
	fmt.Println(calculator.Subtract(20, 10)) // call the Subtract function from the calculator package
}



func number(numbers ...int) int { // numbers is a variadic parameter
	total := 0
	for _, value := range numbers {
		total += value // sum all numbers
	}
	return total // return the total
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

func Rage() {
	// numbers := []int{1000, 100, 10, 1, 0, 11, 111, 1111} // define a slice with 10 elements
	// for i, value := range numbers{
	// 	fmt.Println("Index = ", i, "value =", value)
	// }
	// for _, value := range numbers { // use underscore to ignore the index
	// 	fmt.Println("Value =", value)
	// }
	language := map[string]string{"TH": "Thai", "EN": "English", "JP": "Japanese"}
	for key, value := range language {
		fmt.Println("Key =", key, "Value =", value)
	}
}

// func loop() {
// 	for count := 10; count >= 1; count-- {
// 		if count == 4 {
// 			continue
// 		}
// 		fmt.Println(count)

// 	}
// } // loop 3 times first start from 1 secound < 3 third +1 till 3

func map1() {
	country := map[string]string{"TH": "Thailand", "JP": "Japan", "EN": "England"}
	fmt.Println("The country is", country["JP"])

	Province := map[string]string{}
	Province["ANK"] = "Akihabara" //key
	Province["TKY"] = "Tokyo"
	Province["BKK"] = "Bangkok"

	value, check := Province["ANK"] // check if the key exists
	if check {
		fmt.Println("The Province is", value)
	} else {
		fmt.Println("The key does not exist")
	}

}

func slice() {
	numbers := []int{28, 05, 2001, 0, 8} // define a slice
	numbers = append(numbers, 100)       // append a value to the slice
	numbers = append(numbers, 10)
	fmt.Println("The number is", numbers)
	fmt.Println("The first number is", numbers[0])
	fmt.Println("This is all numbers", numbers[:]) // print all numbers
	fmt.Println(numbers[1:])                       // index 1 to the end
	fmt.Println(numbers[1:3])                      // index 1 to 2
	fmt.Println(numbers[:3])                       // index 0 to 2
}

func array() {
	// number1 := 100
	// number2 := 200
	// number3 := 300
	// var numbers [3]int = [3]int{28, 05, 2001}
	numberss := [...]int{28, 05, 2001, 0, 8} // [...] means the size of the array is determined by the number of elements

	fmt.Println("The first number is", numberss[0])
	fmt.Println("The second number is", numberss[1])
	fmt.Println("The third number is", numberss[2])
	size := len(numberss) // length of array
	fmt.Println("The size of the array is", size)
}

func toswitch() {
	var name string
	fmt.Print("Enter your name account: ")
	fmt.Scanf("%s", &name) // %s is a format specifier for string
	fmt.Println("Hello", name)

	var todo int
	fmt.Println("Enter 1 to deposit, 2 to withdraw")
	fmt.Scanf("%d", &todo) // %d is a format specifier for int

	switch todo {
	case 1:
		var deposit int
		fmt.Print("Enter amount to deposit: ")
		fmt.Scanf("%d", &deposit)
		fmt.Println("You deposited", deposit)
	case 2:
		var withdraw int
		fmt.Print("Enter amount to withdraw: ")
		fmt.Scanf("%d", &withdraw)
		fmt.Println("You withdrew", withdraw)
	default:
		fmt.Println("Invalid option")
	}
}

func bank() {
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

func math1() {
	var num1 int = 10
	var num2 int = 20
	fmt.Println("Result:", num1+num2)
	fmt.Println("Result:", num1-num2)
	fmt.Println("Result:", num1*num2)
	fmt.Println("Result:", num1/num2)
	fmt.Println("Result:", num1%num2)
}

func math2() {
	number1, number2 := 10, 5
	fmt.Println("Is this equal:", number1 == number2)
	fmt.Println("Is this not equal:", number1 != number2)
	fmt.Println("Is this greater than:", number1 > number2)
}

func profile() {

	// var studentName string = "Onwanya"
	// var studentAge int = 23
	// var studentGrade float32 = 4.5
	// var studentResult bool = true
	// fmt.Println("Hello everyone, My name is", studentName)

	// fmt.Println("I am", studentAge, "years old")
	// fmt.Println("Nice to meet you all, I'm in grade", studentGrade)
	// fmt.Println("I am a good student:", studentResult)

	name := "Jane"
	age := 23
	// grade := 4.5
	// result := true
	fmt.Printf("My name is %v\n", name)
	fmt.Printf("Type of name is %T\n", name) // show type of variable
	fmt.Printf("Age = %v", age)
	// fmt.Println("and Hi everyone, you can call me", name)
	// fmt.Println("I am", age, "years old")
	// fmt.Println("I'm", grade)
	// fmt.Println("I am a good student:", result)
}
