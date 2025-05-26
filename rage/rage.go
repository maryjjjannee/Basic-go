package main

import "fmt"

func main() {
	
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
