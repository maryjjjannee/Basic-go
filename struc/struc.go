package main

import "fmt"

type Product struct {
	name string
	price float64
	category string
	discount int
}

func main() {
	product1 := Product{name: "Laptop", price: 50000, category: "Electronics", discount: 10}
	product2 := Product{name: "Mouse", price: 1000, category: "Electronics", discount: 20}
	fmt.Println(product1)
	fmt.Println(product2)
}