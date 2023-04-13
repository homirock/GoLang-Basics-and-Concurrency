package main

import (
	"fmt"
	"math"
)

type Geometry interface {
	area() float64
	per() float64
}

type circle struct {
	Radius float64
}

type rect struct {
	Height, Width float64
}

func (c *circle) area() float64 {
	return math.Pi * c.Radius * c.Radius
}
func (c *circle) per() float64 {
	return 2 * math.Pi * c.Radius
}

func (d *rect) area() float64 {
	return d.Height * d.Width
}

func (d *rect) per() float64 {
	return 2 * (d.Height + d.Width)
}

func measurement(m Geometry) {
	fmt.Println(m.area())
}

func main() {

	rect := &rect{Height: 1, Width: 2}
	measurement(rect)

}
