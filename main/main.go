package main

import "fmt"

func main() {
	// messages := []string{
	// 	"message 1",
	// 	"message 2",
	// 	"message 3",
	// 	"message 4",
	// }

	// for i, message := range messages { // Аналог for i := 0; i < len(messages); i++
	// 	fmt.Println(messages[i])
	// 	fmt.Println(message) // Аналог fmt.Println(messages[i])
	// }

	counter := 0
	for {
		if counter == 100 {
			break // Оператор break полностью выходит из цикла
		}
		counter++
		fmt.Println(counter)
	}
}
