package main

import "fmt"

var message string

func init() {
	message = "new value"
}

func main() {
	fmt.Println(message)
}
