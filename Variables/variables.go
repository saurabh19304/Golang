package main

import "fmt"

func main() {

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int //not initialized so the value is 0
	fmt.Println(e)

	f := "apple" //shorthand to declare a variable ans initializw a variable this is only available inside the function
	fmt.Println(f)
}
