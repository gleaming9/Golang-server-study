package errorExample

import (
	"fmt"
)

/*
fmt.Errorf는 포맷팅 기능을 제공하여, 오류 메시지에 동적 데이터를 포함할 수 있다.
%f는 부동 소수점 숫자를 포맷팅하는 방식으로, a의 값을 오류 메시지에 포함시키는 데 사용됨.
오류 메시지가 고정된 텍스트가 아니라, 상황에 맞는 구체적인 정보를 포함한 메시지를 생성할 수 있음.
*/

func divide2(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide %f by zero", a)
	}
	return a / b, nil
}

func DynamicError() {
	result, err := divide2(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}
