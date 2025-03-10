package shape

import (
	"math"
)

type Shape interface {
	ShapeWithArea
	ShapeWithPerimeter
}

// Интерфейс для вычисления площади
type ShapeWithArea interface {
	Area() float32
}

// Интерфейс для вычисления периметра
type ShapeWithPerimeter interface {
	Perimeter() float32
}

// Определяем структуру Square, представляющую квадрат
type Square struct {
	sideLength float32
}

func NewSquare(length float32) Square {
	return Square{sideLength: length}
}

// Реализуем метод Area() для структуры Square
func (s Square) Area() float32 {
	return s.sideLength * s.sideLength
}

// Реализуем метод Perimeter() для структуры Square
func (s Square) Perimeter() float32 {
	return s.sideLength * 4
}

// Определяем структуру Circle, представляющую круг
type Circle struct {
	radius float32
}

// Реализуем метод Area() для структуры Circle
func (c Circle) Area() float32 {
	return c.radius * c.radius * math.Pi
}

// Реализуем метод Perimeter() для структуры Circle
func (c Circle) Perimeter() float32 {
	return 2 * c.radius * math.Pi
}
