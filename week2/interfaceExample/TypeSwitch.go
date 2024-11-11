package interfaceExample

import "fmt"

// 타입 스위치를 사용하여 여러 타입 처리
func Describe3(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("This is an integer:", v)
	case string:
		fmt.Println("This is a string:", v)
	default:
		fmt.Println("Unknown type.")
	}
}

func TypeSwitch() {
	Describe3(42)
	Describe3("Go Programming")
	Describe3(3.14)
}
