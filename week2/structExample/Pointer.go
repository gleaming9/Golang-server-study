package structExample

import "fmt"

// 포인터 리시버를 사용한 메서드
func (p *Person) HaveBirthday() {
	p.Age++
}

func Pointer() {
	p := Person{Name: "Charlie", Age: 28}

	fmt.Println("Before birthday:", p)
	p.HaveBirthday() // 나이를 증가시키는 메서드 호출
	fmt.Println("After birthday:", p)
}
