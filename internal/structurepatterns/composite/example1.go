package composite

import "fmt"

// Component — общий интерфейс для файлов и директорий.
type Component1 interface {
	Print(string)
	GetSize() int
}

type File struct {
	Name string
	Size int
}

func (f *File) Print(prefix string) {
	fmt.Println(prefix+f.Name, f.Size)
}

func (f *File) GetSize() int {
	return f.Size
}

type Dir struct {
	Name     string
	Children []Component1
}

// Print печатает имя директории, её размер и содержимое.
func (d *Dir) Print(prefix string) {
	fmt.Println(prefix+d.Name, d.GetSize())
	for _, v := range d.Children {
		v.Print(prefix + "  ")
	}
}

// GetSize возвращает общий размер всех файлов в директории.
func (d *Dir) GetSize() int {
	var sum int
	for _, v := range d.Children {
		sum += v.GetSize()
	}
	return sum
}

func Example1() {
	root := &Dir{
		Children: []Component1{
			&File{Name: "file1", Size: 778},
			&File{Name: "file2", Size: 222},
			&Dir{
				Children: []Component1{
					&File{Name: "file3", Size: 64},
					&File{Name: "file4", Size: 36},
				},
				Name: "subfolder",
			},
		},
		Name: "root",
	}
	root.Print("")
}
