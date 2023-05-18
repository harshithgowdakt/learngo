package ds

import "fmt"

func RunMapExample() {

	m := make(map[string]int)

	m["harshith"] = 1
	m["gowda"] = 2

	fmt.Println("Map ::", m)

	fmt.Println("harshith ::", m["harshith"])
	fmt.Println("gowda ::", m["gowda"])
	fmt.Println("0", m["ramdom"])
}
