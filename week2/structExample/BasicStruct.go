package structExample

import "fmt"

type Person struct {
	Name string
	Age  int
}

func BasicStruct() {
	p := Person{Name: "John", Age: 30}
	fmt.Println(p)

	fmt.Println("Name : ", p.Name)
	fmt.Println("Age : ", p.Age)
}
