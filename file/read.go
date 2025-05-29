package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("/Users/mohara/GO/file/Test.csv")
	if err != nil {  // nil mean null, ig no err == nil
		panic(err)
	}

	scanner := bufio.NewScanner(file) // bufio
	count := 1 
	for scanner.Scan(){
		line := scanner.Text()
		item := strings.Split(line, ",") 
		// fmt.Printf("Line %d : %s\n", count, line)  // %d is for integer, %s is for string
		fmt.Println(item[0]) // print the line and the item
		count++ // count the line +1
	}
}