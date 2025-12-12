package main

import (
	"fmt"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r *Rectangle) Scale(factorX float64, factorY float64) {
	r.Width *= factorX
	r.Height *= factorY
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

func CalculateArea(s Shape) float64 {
	return s.Area()
}

func CalculatePerimeter(s Shape) float64 {
	return s.Perimeter()
}

func GuessType(s Shape) {
	// switch s.(type) {
	// case Rectangle:
	// 	fmt.Println("\nIt's a Rectangle")
	// case Circle:
	// 	fmt.Println("\nIt's a Circle")
	// }

	// much wierd way to guess, just to learn type assertions
	_, isRectangle := s.(Rectangle)
	_, isCircle := s.(Circle)

	if isRectangle {
		fmt.Println("\nIt's a Rectangle")
	} else if isCircle {
		fmt.Println("\nIt's a Circle")
	} else {
		fmt.Println("\nUnknown Type!")
	}
}

func main() {
	var rect Rectangle = Rectangle{10, 5}
	var circle Circle = Circle{
		Radius: 7,
	}

	rect.Scale(2, 4) // Scaling rectangle dimensions

	fmt.Printf(
		"Area of rectangle is %v and Perimeter is %v",
		CalculateArea(rect),
		CalculatePerimeter(rect),
	)

	fmt.Printf(
		"\nArea of circle is %v and Perimeter is %v",
		CalculateArea(circle),
		CalculatePerimeter(circle),
	)

	// Type guessing
	GuessType(rect)
	GuessType(circle)
}

/*
  Things to practice:
   - Embedding interfaces & structs
   - Multiple interfaces Implementation
   - Empty interfaces
*/
