package main

import (
	"fmt"
	"math"
)

type Shape interface {
	ShapeWithArea
	ShapeWithPerimeter
}

type ShapeWithArea interface {
	Area() float32
}

type ShapeWithPerimeter interface {
	Perimeter() float32
}

type Square struct {
	sideLenght float32
}

func (s Square) Area() float32 {
	return s.sideLenght * s.sideLenght
}

func (s Square) Perimeter() float32 {
	return s.sideLenght * 4
}

type Circle struct {
	radius float32
}

func (c Circle) Area() float32 {
	return c.radius * c.radius * math.Pi
}

func main() {
	square := Square{5}
	// circle := Circle{8}

	printShapeArea(square)
	// printShapeArea(circle) Ошибка, так как ...

}

func printShapeArea(shape Shape) {
	fmt.Println(shape.Area())
	fmt.Println(shape.Perimeter())
}
