package builder

import "fmt"

// Object — объект с параметром.
type Object struct {
	// данные объекта
	// ...
	// настраиваемые поля объекта
	Mode int
	Path string
}

// SetMode — пример функции, которая присваивает поле Mode.
func (o *Object) SetMode(mode int) *Object {
	o.Mode = mode
	return o
}

// SetPath — пример функции, которая присваивает поле Path.
func (o *Object) SetPath(path string) *Object {
	o.Path = path
	return o
}

// WithMode — пример функции, которая присваивает поле Mode.
func WithMode(mode int) func(*Object) {
	return func(o *Object) {
		o.Mode = mode
	}
}

// WithPath — пример функции, которая присваивает поле Path.
func WithPath(path string) func(*Object) {
	return func(o *Object) {
		o.Path = path
	}
}

// NewObject — функция-конструктор объекта.
func NewObject(opts ...func(*Object)) *Object {
	o := &Object{}

	// вызываем все указанные функции для установки параметров
	for _, opt := range opts {
		opt(o)
	}
	return o
}

func Example2() {
	o := NewObject(WithMode(10), WithPath(`root`))
	fmt.Println(o)
}
