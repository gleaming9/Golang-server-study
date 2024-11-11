package interfaceExample

import "fmt"

// Speaker 인터페이스 정의
type Speaker interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct{}

func (c Cat) Speak() string {
	return "Meow!"
}

// 인터페이스를 매개변수로 사용하는 함수
func MakeSound(s Speaker) {
	fmt.Println(s.Speak())
}

func BasicInterface() {
	var s Speaker
	d := Dog{}
	c := Cat{}

	s = d
	fmt.Println(s.Speak())
	s = c
	fmt.Println(s.Speak())

	MakeSound(d)
	MakeSound(c)
}
