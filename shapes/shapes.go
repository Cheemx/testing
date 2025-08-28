// shapes provide functions to provide perimeters and areas of different shapes!
package shapes

import "math"

// Rectangle type supports representing rectangle shape in geometry.
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

// Circle type supports representing circle shape in geometry.
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Shape interface {
	Area() float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

// Permieter provides perimeter of the shape.
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Since Go doesn't support method overloading so
// following code will provide compile time error

// func Area(circle Circle)  {
// 	return 0.0
// }

// Area redeclraed in this block

// S rather than that we use Methods!

// Area provides surface area of shape
func Area(rectangle Rectangle) float64 {
	return rectangle.Height * rectangle.Width
}
