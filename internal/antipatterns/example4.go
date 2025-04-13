package antipatterns

import (
	"fmt"
	"io"
)

func equal(x, y interface{}) bool {
	return x == y
}

type Int int

func Example4_bad() {
	var s fmt.Stringer
	var r io.Reader
	fmt.Println(equal(s, r)) // true
	var i int = 5
	var I Int = 5
	fmt.Println(equal(i, I)) // false
	var sl1, sl2 []int
	fmt.Println(equal(sl1, sl2)) // panic
}
