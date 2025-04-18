package iterator

import "fmt"

// Iterator — интерфейс для получения следующего элемента.
type Iterator interface {
	Next() (string, bool)
}

// Exported — тип, реализующий интерфейс Iterator.
type Exported struct {
	ID        string
	invisible []string
	cursor    int
}

func NewExported(s ...string) *Exported {
	e := new(Exported)
	e.invisible = append(e.invisible, s...)
	return e
}

// Next — метод, реализующий шаблон Итератор.
func (e *Exported) Next() (string, bool) {
	if e.cursor <= len(e.invisible) {
		e.cursor++
	}
	return e.invisible[e.cursor-1], e.cursor < len(e.invisible)
}

func Example2() {
	// клиентский код
	e := NewExported("foo", "bar", "buzz", "oups")
	for {
		next, hasNext := e.Next()
		fmt.Println(next)
		if !hasNext {
			break
		}
	}
}
