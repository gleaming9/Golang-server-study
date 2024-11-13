package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

/*
rw: 실제 HTTP 응답을 보내는 http.ResponseWriter입니다. 응답 본문과 헤더를 작성할 수 있습니다.
mw: io.MultiWriter로 생성된 io.Writer 타입의 필드로, 응답 본문을 동시에 두 곳(실제 응답과 버퍼)에 기록합니다.
status: 응답의 상태 코드(예: 200, 404 등)를 저장하는 필드입니다.
상태 코드가 설정되지 않은 경우 기본값으로 StatusOK(200)이 설정됩니다.
*/
type rwWrapper struct {
	rw     http.ResponseWriter
	mw     io.Writer
	status int
}

/*
NewRwWrapper 는 rwWrapper 인스턴스를 생성하는 함수입니다.
io.MultiWriter(rw, buf)를 통해 응답 본문이 클라이언트와 버퍼(로그에 저장될) 두 곳에 동시에 기록될 수 있도록 설정합니다.
이를 통해 클라이언트로 보내진 응답이 버퍼에 그대로 남아 있어 로그로 기록할 수 있습니다.
*/
func NewRwWrapper(rw http.ResponseWriter, buf io.Writer) *rwWrapper {
	return &rwWrapper{rw: rw, mw: io.MultiWriter(rw, buf)}
}

// Header : http 응답 헤더 반환
func (r *rwWrapper) Header() http.Header {
	return r.rw.Header()
}

// Write() 메서드는 HTTP 응답 본문을 작성합니다.
func (r *rwWrapper) Write(i []byte) (int, error) {
	if r.status == 0 {
		r.status = http.StatusOK
	}
	return r.mw.Write(i)
}

func (r *rwWrapper) WriteHeader(statusCode int) {
	r.status = statusCode
	r.rw.WriteHeader(statusCode)
}

/*
NewLogger 미들웨어는 요청을 받으면 rwWrapper를 통해 응답을 가로챕니다.
응답 본문과 상태 코드가 rwWrapper에 기록되면서 버퍼에도 저장됩니다.
응답이 완료된 후, 로그에 클라이언트 IP, 요청 메서드, 요청 경로, 상태 코드, 응답 본문을 기록합니다.
이 미들웨어를 통해 모든 HTTP 요청에 대한 응답 정보를 로그에 남길 수 있습니다.
*/
func NewLogger(l *log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			buf := &bytes.Buffer{}
			rww := NewRwWrapper(w, buf)
			next.ServeHTTP(rww, r)
			l.Printf("%s %s %s %d %s", r.RemoteAddr, r.Method, r.URL, rww.status, buf)
		})
	}
}
