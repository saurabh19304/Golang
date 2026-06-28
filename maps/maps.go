package maps

import (
	"fmt"
	"maps"
)

func main() {

	m := make(map[string]int)

	m["k1"] = 2
	m["k2"] = 5

	fmt.Println("map:", m)
}
