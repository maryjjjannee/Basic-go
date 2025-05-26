package main

import "fmt"

func main() {
	
	for count := 10; count >= 1; count-- {
		if count == 4 {
			continue
		}
		fmt.Println(count)

	}
} // loop 3 times first start from 1 secound < 3 third +1 till 3
