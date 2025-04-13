package singleton

import (
	"log"
	"os"
	"sync"
)

//var dumpFile *os.File
//
//func Dump(data []byte) error {
//	if dumpFile == nil {
//		// создаём файл при первом обращении
//		fname, err := os.Executable()
//		if err == nil {
//			dumpFile, err = os.Create(fname + `.dump`)
//		}
//		if err != nil {
//			panic(err)
//		}
//	}
//	_, err := dumpFile.Write(data)
//	return err
//}

/*
Задание 1 из 7
Горутины в программе могут вызывать функцию для сохранения дампа данных в файл.
Если во время работы программы сохранения не происходит, то файл не должен создаваться.
Разработчик предложил решение, которое может работать неверно при одновременном вызове функции Dump.
Исправьте этот фрагмент кода.
*/

var (
	dumpFile *os.File
	once     sync.Once
)

func Task1(data []byte) error {
	once.Do(
		func() {
			fname, err := os.Executable()
			if err == nil {
				dumpFile, err = os.Create(fname + `.dump`)
				log.Println(fname + `.dump`)
			}
			if err != nil {
				panic(err)
			}
		})
	_, err := dumpFile.Write(data)
	return err
}
