package antipatterns

import "net/http"

type httpFunc func(r *http.Request)

var Funcs = make(map[string]httpFunc) // пустой интерфейс может принять любое значение

func DBInsert(r *http.Request) {
	// логика вставки
}

func DBDelete(r *http.Request) {
	// логика удаления
}

func Task1() {
	Funcs["DBInsert"] = DBInsert
	Funcs["DBDelete"] = DBDelete
	Funcs["DBChange"] = func(r *http.Request) {
		// логика изменения
	}
	r := new(http.Request)
	Funcs["DBInsert"](r)
	Funcs["DBInsert"](new(http.Request))
	// ...
}
