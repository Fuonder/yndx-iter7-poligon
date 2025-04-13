package antipatterns

import "math/rand"

//
//import (
//	"fmt"
//	"math/rand"
//	"sync"
//	"time"
//)
//
//const IDLength = 6
//// декларируем алфавит символов для идентификатора
//var Alphabet = []rune("0123456789abcdefghijklmnopqrstuvwxyz")
//
//// GenerateName создаёт случайный идентификатор.
//func GenerateName() string {
//	// аллокация слайса
//	name := make([]rune, IDLength)
//	// заполняем слайс псевдослучайными символами из алфавита
//	for i := 0; i < IDLength; i++ {
//		name[i] = Alphabet[rand.Intn(len(Alphabet))]
//	}
//	// приводим слайс к типу string и возвращаем
//	return string(name)
//}
//
//// хранилище идентификатора
//type Generator struct {
//	Names map[string]struct{}
//	// защищено мьютексом для безопасного конкурентного доступа
//	Mutex sync.Mutex
//}
//
//// конструктор хранилища
//func NewGenerator() *Generator {
//	return &Generator{
//		Names: make(map[string]struct{}),
//	}
//}
//
//// метод хранилища, возвращающий уникальный идентификатор
//func (gen *Generator) GetUniqueName() string {
//	for {
//		// генерируем идентификатор
//		name := GenerateName()
//		// закрываем замок мьютекса
//		gen.Mutex.Lock()
//		// проверяем идентификатор на уникальность
//		if _, ok := gen.Names[name]; !ok {
//			// если идентификатор не регистрировался в хранилище,
//			// то регистрируем
//			gen.Names[name] = struct{}{}
//			// открываем замок
//			gen.Mutex.Unlock()
//			// и возвращаем
//			return name
//		}
//		// открываем замок
//		gen.Mutex.Unlock()
//	}
//}
//
//func Example1_bad() {
//	// gen содержит сгенерированные случайные идентификаторы
//	gen := NewGenerator()
//	// используем генератор конкурентно
//	// из нескольких горутин
//	for i := 0; i < 5; i++ {
//		go func() {
//			for j := 0; j < 10; j++ {
//				fmt.Println(gen.GetUniqueName())
//			}
//		}()
//	}
//	time.Sleep(1 * time.Second)
//}

const IDLength = 6

// декларируем алфавит символов для идентификатора
var Alphabet = []rune("0123456789abcdefghijklmnopqrstuvwxyz")

type Generator struct {
	Names map[string]struct{}
	// снабжаем генератор не мьютексом,
	// а каналом передачи сообщений
	Ch chan string
}

// конструктор генератора
func NewGenerator() *Generator {
	gen := &Generator{
		Names: make(map[string]struct{}),
		Ch:    make(chan string),
	}
	// вся работа с генератором будет выполняться
	// в одной специально запущенной горутине
	go gen.GenProcess()
	return gen
}

func GenerateName() string {
	// аллокация слайса
	name := make([]rune, IDLength)
	// заполняем слайс псевдослучайными символами из алфавита
	for i := 0; i < IDLength; i++ {
		name[i] = Alphabet[rand.Intn(len(Alphabet))]
	}
	// приводим слайс к типу string и возвращаем
	return string(name)
}

// метод генерации,
// только он имеет доступ к мапе-хранилищу
// запускается в отдельной горутине
// остальным потокам отдаёт идентификаторы через канал
func (gen *Generator) GenProcess() {
	for {
		name := GenerateName()
		if _, ok := gen.Names[name]; !ok {
			gen.Names[name] = struct{}{}
			gen.Ch <- name
		}
	}

}

// метод для внешних вызовов генератора,
// его можно безопасно вызывать
// из нескольких горутин
func (gen *Generator) GetUniqueName() string {
	return <-gen.Ch
}
