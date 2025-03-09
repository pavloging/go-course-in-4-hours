package main

import "fmt"

type Age int
// С помощью чего можно прикреплять к ним методы
func (a Age) isAdult() bool {
    return a >= 18
}

type User struct {
	name   string
	age    Age
	sex    string
	weight int
	height int
}
 
func NewUser(name, sex string, age, weight, height int) User {
	return User{
		name:   name,
		sex:    sex,
		age:    Age(age), // Используем приведение типа
		weight: weight,
		height: height,
	}
}

func main() {
	user1 := NewUser("Vova", "Male", 11, 75, 188)

    fmt.Println(user1.age.isAdult())
}