package errorExample

import "fmt"

// 사용자 정의 오류 타입 정의
type DivideError struct {
	Dividend float64
	Divisor  float64
}

/*
Go에서 사용자 정의 오류 타입을 만들 때, 반드시 Error()라는 이름의 메서드를 정의해야만 해당 타입이 error 인터페이스를 구현하게 됨.
error 인터페이스는 아래와 같이 정의돼 있음

	type error interface {
	    Error() string
	}

따라서, Error()라는 이름의 메서드를 구현하지 않으면 그 타입은 error 인터페이스를 만족하지 않게 됨.
만약 Error() 메서드 대신 다른 이름을 사용하면, 해당 타입은 더 이상 error 인터페이스를 구현한 것이 아니기 때문에,
error 타입으로 인식되지 않아 오류가 발생할 수 있다.
오버라이딩 개념
*/
func (e *DivideError) Error() string {
	return fmt.Sprintf("cannot divide %f by %f", e.Dividend, e.Divisor)
}

// 나눗셈 함수
func divide3(a, b float64) (float64, error) {
	if b == 0 {
		return 0, &DivideError{Dividend: a, Divisor: b}
	}
	return a / b, nil
}

func CustomError() {
	result, err := divide3(10, 0)
	if err != nil {
		// 사용자 정의 오류를 출력
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}
