package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}
type rectangle struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}
func (r rectangle) perimeter() float64 {
	return (r.width + r.height) * 2
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func main() {
	r := rectangle{height: 2, width: 3}
	c := circle{radius: 2}
	measure(r)
	measure(c)
}
