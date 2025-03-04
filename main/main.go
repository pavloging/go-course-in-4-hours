package main

import "fmt"

func main() {
	messages := make([]string, 5) // len=5, cap=5
	messages = append(messages, "6")

	fmt.Println(len(messages))
	fmt.Println(cap(messages))

	fmt.Println(messages)

}
