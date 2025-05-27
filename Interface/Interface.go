package main

import (
	"fmt"
	"math"
)

type geometry interface { // geometry is an interface
	area() float64
	perim() float64
}

type react struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r react) area() float64{
	return r.width * r.height // area of rectangle = width * height
}

func (r react) perim() float64{
	return 2*r.width + 2*r.height // perimeter of rectangle = 2(width + height)
}

func (c circle) area() float64{
	return math.Pi * c.radius * c.radius // area of circle = πr²
}

func (c circle) perim() float64{
	return  2*math.Pi * c.radius // perimeter of circle = 2πr
}

func measure(g geometry) {
     fmt.Println(g)
	 fmt.Println("Area = ", g.area())
	 fmt.Println("Perim = ", g.perim())  
}

func main() {
	r := react{width: 5, height:10}
	c := circle{radius: 15}

	measure(r) 
	measure(c) 
}