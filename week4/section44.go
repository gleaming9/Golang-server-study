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
이 반환된 함수는 next http.Handler를 매개변수로 받아 http.Handler를 반환합니다.
next는 이후에 실행될 핸들러를 가리킵니다. 이 핸들러는 이 미들웨어 이후에 실행됩니다.

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
