package main

import "fmt"

func main() {

	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	b := [5]int{0, 1, 2, 3, 4}
	fmt.Println("ar", b)

	b = [...]int{0, 1, 2, 3, 4}
	fmt.Println("len", b)

	b = [...]int{100, 3: 200, 4}
	fmt.Println("ind", b)
}
