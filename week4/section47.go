package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

/*
요청 바디를 로그에 남기는 미들웨어

본문 읽기: 요청 본문을 io.ReadAll 로 읽어 body 에 저장합니다.

에러 처리: 본문을 읽는 중에 에러가 발생하면 이를 로그로 기록하고, 클라이언트에 400 Bad Request 응답을 반환합니다.

본문 복구: 본문을 읽은 후 r.Body 를 새로운 버퍼로 대체하여 이후 핸들러에서도 본문을 다시 사용할 수 있도록 합니다.
요청 본문은 한 번 읽으면 EOF(End of File) 상태가 되어 다시 읽을 수 없습니다. 따라서 본문을 읽은 후,
이를 새로운 io.Reader 로 복구해 r.Body 에 다시 할당합니다.
bytes.NewBuffer(body)는 body 바이트 배열로부터 새로운 버퍼를 생성하여 이를 r.Body 에 할당합니다.
io.NopCloser 는 io.Reader 를 받아 이를 닫을 수 있는 인터페이스(io.ReadCloser)로 변환해줍니다.

핸들러 실행: 이후 핸들러에 요청을 전달하여 실제로 요청을 처리합니다.

이런 식으로 구현하는 것은 성능에 영향을 주므로 주의해야 합니다.
미들웨어 구현은 적용한 엔드포인트가 실행될 때마다 반드시 실행된다는 문제가 있습니다.
요청 바디를 매번 복사하는 미들웨어를 사용할 때는 미들웨어 적용 전후로 성능에 영향이 없는지 확인.

예시:
GET /users : 모든 사용자 목록을 반환하는 엔드포인트
POST /users : 새로운 사용자를 생성하는 엔드포인트
GET /users/{id} : 특정 사용자의 정보를 반환하는 엔드포인트
DELETE /users/{id} : 특정 사용자를 삭제하는 엔드포인트

엔드포인트는 서버가 특정 작업을 수행하기 위해 외부에서 요청을 받을 수 있는 특정 URL이며,
미들웨어는 이 요청이 엔드포인트에서 처리되기 전후에 항상 실행되는 코드입니다.
*/

func RequestBodyLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("Failed to log request body: %v", err)
			http.Error(w, "Failed to get request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		r.Body = io.NopCloser(bytes.NewBuffer(body))
		next.ServeHTTP(w, r)
	})
}
