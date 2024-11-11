package errorExample

import "fmt"

// 패닉을 발생시키고 이를 복구하는 예제
func riskyOperation() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println("Starting risky operation...")
	panic("Something went wrong!") // 패닉 발생
	fmt.Println("This line will not be executed.")
}

func PanicRecover() {
	riskyOperation()
	fmt.Println("Program continues after recovery.")
}
