package main

import "fmt"

type User struct {
	name   string
	age    int
	sex    string
	weight int
	height int
}

func NewUser(name, sex string, age, weight, height int) User {
	return User{
		name:   name,
		sex:    sex,
		age:    age,
		weight: weight,
		height: height,
	}
}

func main() {
	user1 := NewUser("Vova", "Male", 23, 75, 188)
	fmt.Printf("%+v", user1)
}
