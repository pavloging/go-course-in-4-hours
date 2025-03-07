package main

import "fmt"

func main() {
	// Map предстовляет собой тип данных который является словарем
	// Работает по концепции ключ значение
	// Ключ в map должен быть уникальным (не должен дублироваться)
	users := map[string]int{
		"Alex":  10,
		"Ivan":  25,
		"Artem": 42,
	}

	// Перезапись
	users["Alex"] = 52

	// Цикл по map
	for key, value := range users {
		fmt.Println(key, value)
	}

	// Удаление элемента map
	delete(users, "Ivan")

	// Если у нас не инициализировнна map, то мы в неё не можем писать
	var users1 map[string]int
	users1["Vova"] = 15
}
