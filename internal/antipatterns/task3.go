package antipatterns

import (
	"io"
	"os"
)

//func ConditionalWrite(w io.Writer) (int, error) {
//	// используем рефлексию для определения конкретного типа
//	if wtype := reflect.TypeOf(w); wtype.String() == "*os.File" {
//		// если это *os.File, предпринимаем действия
//	}
//	// действуем по-другому
//}

//func assertionalWrite(w io.Writer) (int, error) {
//	switch w.(type) {
//	case *os.File:
//	default:
//	}
//
//	return 0, nil
//}

func ConditionalWrite(w io.Writer) (int, error) {
	// используем type assertion
	if f, ok := w.(*os.File); ok {
		// если это *os.File предпринимаем действия
		f.Stat()
		return 0, nil
	}
	// действуем по-другому
	return 0, nil
}
