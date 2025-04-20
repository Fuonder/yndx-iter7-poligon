package formatting

import (
	"fmt"
)

func add(a, b int) int {
	return a + b
}

func Example1() {
	a := 100 // левое слагаемое
	b := -1  // правое слагаемое
	fmt.Printf("Сумма: %d", add(a, b))
}
