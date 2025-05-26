package main
import "fmt"
func main() {
	
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