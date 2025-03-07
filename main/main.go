package main

import "fmt"

func main() {
	defer handlePanic()

	messages := []string{
		"message 1",
		"message 2",
		"message 3",
		"message 4",
	}

	messages[4] = "message 5" // Получаем панику и выходим из main()

	fmt.Println(messages)
}

func handlePanic() {
	if r := recover(); r != nil {
		fmt.Println(r)
	}

	fmt.Println("handlePanic() выполнилась успешно")
}
