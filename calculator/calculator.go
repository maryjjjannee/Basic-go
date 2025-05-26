package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin) // get input from bufio

func getInput(prompt string) float64 {
	fmt.Printf("%v", prompt)                                      // %v output datatype
	input, _ := reader.ReadString('\n')                            // read input until newline
	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64) // convert string to float64
	if err != nil {                                                // nil check for error
		message := fmt.Sprintf("%v must number only", prompt)
		panic(message) // panic same as print
	}
	return value
}

func add(value1, value2 float64) float64 {
	return value1 + value2
}

func subtract(value1, value2 float64) float64 {
	return value1 - value2
}

func multiply(value1, value2 float64) float64 {
	return value1 * value2
}

func divide(value1, value2 float64) float64 {
	return value1 / value2
}

func getOperator() string { // sreing คือ ชนิดข้อมูลที่ return
	fmt.Println("operator is (+ - * /): ")
	op, _ := reader.ReadString('\n')
	return strings.TrimSpace((op)) // trim space from input
}

func main() {
	value1 := getInput("value1 =")
	value2 := getInput("value2 =")

	var result float64

	switch operator := getOperator(); operator {
	case "+":
		result = add(value1, value2)
	case "-":
		result = subtract(value1, value2)
	case "*":
		result = multiply(value1, value2)
	case "/":
		result = divide(value1, value2)
	default:
		panic("wrong operator " + operator)
	}
	fmt.Printf("result is %v", result)

}
