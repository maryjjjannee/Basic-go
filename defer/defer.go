package main

import "fmt"

func add(a, b float64) {

	result := a + b
	fmt.Println("Result:", result)
}

//defer is make sure that the function is executed at the end of the function

func loop() {
	for i := 0; i<10 ; i++ {
		 fmt.Println("Looping:", i)  
	}
}

func deferloop(){
	for j := 0; j <10 ; j++{
		defer fmt.Println("j = ", j) //use defer will print last in first out (LIFO)
	}
}

func main() {
	// fmt.Println("Welcome to calculator")
	// defer fmt.Println("Exiting calculator")
	// defer add(20, 10)
	// defer add(15,15)
	// defer add(12, 12)
	loop()
	deferloop()
	// if use many defer, will use last in first out (LIFO) last will be first 
}
