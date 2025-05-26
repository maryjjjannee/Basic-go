package main
import "fmt"

func main() {

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
