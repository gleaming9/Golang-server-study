package structExample

import "fmt"

type Animal struct {
	Name string
}

func (a Animal) Speak() string {
	return "I am " + a.Name
}

// Dog 구조체가 Animal 구조체를 임베딩
type Dog struct {
	Animal // Animal 구조체의 필드와 메서드 상속
	Breed  string
}

func Embedding() {
	d := Dog{
		Animal: Animal{Name: "Buddy"},
		Breed:  "Golden Retriever",
	}

	// Dog 구조체는 Animal의 메서드를 사용할 수 있음
	fmt.Println(d.Speak())
	fmt.Println("Breed:", d.Breed)
}
