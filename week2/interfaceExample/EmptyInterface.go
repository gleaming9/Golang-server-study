package interfaceExample

import "fmt"

// 빈 인터페이스를 매개변수로 받는 함수
func Describe(i interface{}) {
	fmt.Printf("Type: %T, Value: %v\n", i, i)
}

func EmptyInterface() {
	Describe(42)
	Describe("Hello")
	Describe(3.14)
}
