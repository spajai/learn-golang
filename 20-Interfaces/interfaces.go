package main

import "fmt"
import "math"

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	height, width float64
}

type circle struct {
	radius, diameter float64
}

func (r rect) area() float64 {
	return r.height * r.width
}

func (c circle) area() float64 {
	c.diameter = 2 * c.radius
	fmt.Println("circ diameter: ", c.diameter)

	return math.Pi * c.diameter
}

func (r rect) perim() float64 {
	perim := 2 * (r.height + r.height)
	fmt.Println("Peri rect:", perim)
	return perim
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measures(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measures(r)
	measures(c)

}
