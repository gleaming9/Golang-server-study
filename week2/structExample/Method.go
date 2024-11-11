package structExample

import "fmt"

// Person 구조체의 메서드 정의
func (p Person) Greet() string {
	return "my name is " + p.Name + " and I am " + fmt.Sprint(p.Age)
}

func Method() {
	p := Person{Name: "Bob", Age: 30}
	fmt.Println(p.Greet())
}
