package main

import "fmt"

func zerovalue(ivalue int){
	ivalue = 0 
}

func zeropointer(ipointer *int) {
	*ipointer = 0
}

func main() {
	i := 1 
	fmt.Println("i =", i)

	zerovalue(i)
	fmt.Println("i from function zerovalue =", i)

	zeropointer(&i) // & is used to get the address of i
	fmt.Println("i value from function zeropointer =", i)
	fmt.Println("i address from function zeropointer =", &i) // &i is the address of i
}