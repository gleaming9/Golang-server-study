package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
RecoveryMiddleware 는 HTTP 요청을 처리하는 핸들러에 패닉 복구 기능을 추가하는 미들웨어입니다.
요청 처리 중 패닉이 발생하면 recover()로 이를 복구하고, 에러 메시지를 JSON 형식으로 클라이언트에 반환합니다.
이렇게 하면 서버가 중단되지 않고 오류 응답을 일관된 형식으로 전달할 수 있습니다.
*/

func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				jsonBody, _ := json.Marshal(map[string]interface{}{
					"error": fmt.Sprintf("%v", err),
				})
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(jsonBody)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
