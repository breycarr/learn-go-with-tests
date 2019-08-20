package main

import "math"

// Shape is implemented by anything that declares its area
type Shape interface {
	Area() float64
}

// Rectangle represents a rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// Area returns the area of a rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle represents a circle
type Circle struct {
	Radius float64
}

// Area returns the area of a circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Triangle represents a triangle
type Triangle struct {
	Width  float64
	Height float64
}

// Area returns the area of a triangle
func (t Triangle) Area() float64 {
	return (t.Height * t.Width) / 2
}

// Perimeter returns the perimeter of a rectangle
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Area is a deprecated function that was replaced by an Area method
// func Area(rectangle Rectangle) float64 {
// 	return rectangle.Width * rectangle.Height
// }

func main() {

}
