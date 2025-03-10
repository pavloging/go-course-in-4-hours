package main

import (
	"fmt"
	"go-course/shape"
	"time"

	"github.com/zhashkevych/scheduler"
)

func main() {
	t := time.Now()
	fmt.Println(t)

	square := shape.NewSquare(5)
	printShapeArea(square)

	scheduler := scheduler.NewScheduler()
	fmt.Println(scheduler)
}

func printShapeArea(s shape.Shape) {
	fmt.Println(s.Area())
	fmt.Println(s.Perimeter())
}
