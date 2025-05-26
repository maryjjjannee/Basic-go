package main
import "fmt"

func main() {
	math1() // call the math1 function
	math2() // call the math2 function
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