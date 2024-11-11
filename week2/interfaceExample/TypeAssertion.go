package interfaceExample

import "fmt"

func Describe2(i interface{}) {
	// Type Assertion을 사용하여 타입을 단언
	value, ok := i.(int)
	if ok {
		fmt.Println("This is an integer:", value)
	} else {
		fmt.Println("Not an integer.")
	}
}
func TypeAssertion() {
	Describe2(100)
	Describe2("Hello")
}
