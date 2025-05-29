package main

import (
	"fmt"
	"time"
)

func process1(c chan string, data string){ // chan for sent data 
	c <- data // send data to channel

}

func main() { // sent main func to process1
	ch := make(chan string, 1) 
	go process1(ch, "Hello Stitch") // goroutine to process data
	fmt.Println(<-ch) // Receive data from channel
	time.Sleep(5 * time.Second) // Wait for goroutine to finish

}