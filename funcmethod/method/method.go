package main

import (
	"fmt"
	"image/color"
)

type Rect struct {
	length, width, border int
}

func (r Rect) area() int {
	return r.length * r.width
}
func (r Rect) perim() int {
	return (r.length + r.width) * 2
}

type ColorRect struct {
	Rect
	color  color.RGBA
	border int
}

func main() {
	r := Rect{width: 5, length: 10}
	fmt.Println("area: ", r.area())  // area: 50
	fmt.Println("perim:", r.perim()) // perim: 30
	red := color.RGBA{255, 0, 0, 255}
	c := ColorRect{Rect{10, 5, 0}, red, 1}
	c.width += 5
	fmt.Println("area: ", c.area(), c.Rect.border) // area: 100
	fmt.Println("perim:", c.perim(), c.border)     // perim: 40
}
