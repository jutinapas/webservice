package chapters

import "fmt"

const (
	first  = iota
	second = iota
)

const (
	third = iota
	fourth
)

func constants() {
	const x = 3
	fmt.Println(x + 3)
	fmt.Println(x + 1.2)

	const y int = 3
	fmt.Println(y + 3)
	fmt.Println(float32(y) + 1.2)

	fmt.Println(first, second, third, fourth)
}
