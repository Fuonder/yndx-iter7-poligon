package antipatterns

import (
	"fmt"
	"reflect"
	"unsafe"
)

//array := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
//fmt.Println(array) // [0 1 2 3 4 5 6 7 8 9]
//slice := array[:5]
//slice = append(slice, 42)
//fmt.Println(array) // [0 1 2 3 4 42 6 7 8 9]

func Task2() {
	array := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice := array[:5]
	// ограничили длину len слайса, но его вместимость cap не изменилась
	fmt.Println(cap(slice)) // 10
	// получим заголовок слайса
	slHeader := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	// и в нём ограничим capacity
	slHeader.Cap = 5
	// так работает
	fmt.Println(cap(slice)) // 5
	// теперь, поскольку вместительность слайса ограничена,
	// при вызове append ему будет выделен новый участок памяти
	slice = append(slice, 42)
	fmt.Println(slice) // [0 1 2 3 4 42]
	// защитим несущий массив array от неожидаемых изменений
	fmt.Println(array) // [0 1 2 3 4 5 6 7 8 9]

	array2 := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice2 := make([]byte, 5)
	copy(slice2, array2[:5])
	slice2 = append(slice2, 42)
	fmt.Println(slice2)
	fmt.Println(array2)

}
