package mydoc

import "fmt"

func ExampleStudent_SetName() {
	student := Student{}
	student.SetName("John Doe")
	fmt.Println(student.Name)

	// Output:
	// John Doe
}

func ExampleStudent_GetName() {
	student := Student{
		Name: "John Doe",
	}

	fmt.Println(student.GetName())

	// Output:
	// John Doe
}
