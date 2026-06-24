package main

import (
	"fmt"
)

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

func main() {
	var pers1 Person
	var pers2 Person

	pers1.name = "Hege"
	pers1.age = 45
	pers1.job = "Teacher"
	pers1.salary = 6000

	pers2.name = "Cecilie"
	pers2.age = 23
	pers2.job = "Marketing"
	pers2.salary = 5900

	fmt.Println("name:", pers1.name)
	fmt.Println("age", pers1.age)
	fmt.Println("job", pers1.job)
	fmt.Println("salary", pers1.salary)

	fmt.Println("name:", pers2.name)
	fmt.Println("age", pers2.age)
	fmt.Println("job", pers2.job)
	fmt.Println("salary", pers2.salary)
}
