package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	base   float64
	height float64
}

func main() {
	tri := triangle{ base: 5, height: 10}
	sq := square {sideLength: 4}

	printArea(tri)
	printArea(sq)
}

func printArea(sh shape) {
	fmt.Println("Area:", sh.getArea())
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {	
	return 0.5 * t.base * t.height
}
