package main

import (
	"log"
	"net/http"
	"time"
)

/*
MyMiddleware 함수는 HTTP 요청을 처리하는 핸들러를 받아 실행 시간을 측정하는 추가 기능을 더해줍니다.
미들웨어는 요청을 시작할 때 시간을 기록하고, 요청을 처리한 후 경과 시간을 계산합니다.
이 경과 시간을 로깅하여 요청이 얼마나 걸렸는지 쉽게 확인할 수 있습니다.
*/
func MyMiddleware(h http.Handler) http.Handler {
	/*
		http.Handler 인터페이스를 다시 확인해보겠습니다.

		type Handler interface {
		    ServeHTTP(http.ResponseWriter, *http.Request)
		}

		 http.Handler 인터페이스는 ServeHTTP(http.ResponseWriter, *http.Request) 메서드를 가지고 있어야 합니다.
		이 인터페이스를 구현해야만 http.Handler로 사용할 수 있습니다.
		그런데 익명 함수 func(w http.ResponseWriter, r *http.Request) { ... }는 직접적으로 ServeHTTP라는 이름의 메서드를 가지고 있지 않습니다.
		따라서 이 함수는 http.Handler 인터페이스를 만족하지 못합니다.

		 http.HandlerFunc는 func(w http.ResponseWriter, r *http.Request) 형태의 함수를 http.Handler 인터페이스로 변환합니다.
		http.HandlerFunc 타입 자체는 type HandlerFunc func(http.ResponseWriter, *http.Request)로 정의되어 있으며, ServeHTTP 메서드를 구현하고 있습니다.
		따라서 http.HandlerFunc로 래핑된 함수는 자동으로 ServeHTTP 메서드를 제공하게 되어 http.Handler 인터페이스를 만족하게 됩니다.
	*/
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s := time.Now()                                          // 요청 시작 시간 기록
		h.ServeHTTP(w, r)                                        // 요청 처리
		d := time.Now().Sub(s).Milliseconds()                    // 경과 시간 계산
		log.Printf("end %s(%d ms)\n", s.Format(time.RFC3339), d) // 요청 시작 시간과 경과 시간 로깅
	})
}

/*
VersionAdder 미들웨어를 사용하면, App-version 헤더에 버전 정보를 추가한 상태로 다른 핸들러가
요청을 처리할 수 있습니다. VersionAdder 함수는 내부적으로 또 다른 함수를 반환합니다.
이 반환된 함수는 next http.Handler 를 매개변수로 받아 http.Handler 를 반환합니다.
next 는 이후에 실행될 핸들러를 가리킵니다. 이 핸들러는 이 미들웨어 이후에 실행됩니다.

vmw := VersionAdder("1.0.1")
http.Handle("/users", vmw(UsersHandler))
*/
func VersionAdder(v string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r.Header.Add("App-version", v)
			next.ServeHTTP(w, r)
		})
	}
}
