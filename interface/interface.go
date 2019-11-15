package main

import "fmt"

// type Shaper interface {
// 	area() int  // Tính diện tích
// 	perim() int // Tính chu vi
// }

type Rect struct {
	length, width int
}

type Square struct {
	side int
}

func (r Rect) area() int {
	return r.length * r.width
}

func (r Rect) perim() int {
	return (r.length + r.width) * 2
}

func (sq Square) area() int {
	return sq.side * sq.side
}

func (sq Square) perim() int {
	return sq.side * 4
}

type Areaer interface {
	area() int
}

type Perimer interface {
	perim() int
}

type Shaper interface {
	Areaer  // Tính diện tích
	Perimer // Tính chu vi
}

func main() {
	var s Shaper
	r := Rect{length: 5, width: 3}
	s = Shaper(r)
	fmt.Println("Area: ", s.area())       // Area: 15
	fmt.Println("Perimeter: ", s.perim()) // Perimeter: 16

	q := Square{side: 5}
	s = q
	fmt.Println("Area: ", s.area())       // Area: 25
	fmt.Println("Perimeter: ", s.perim()) // Perimeter: 20
}
