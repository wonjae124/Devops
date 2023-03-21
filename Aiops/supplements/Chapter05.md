# 0. webserver with golang, making docker image

```golang
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// 핸들러 로직 작성
	// http.ResponseWriter는 HTTP response에 무언가를 쓸 수 있게한다. Hello docker를 반환하게 한다.
	// http.Request는 입력된 HTTP request를 검토한다
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		log.Println("received request")
		fmt.Fprintf(w, "Hello Docker!!")
	})

	log.Println("Start server")
	server:= &http.Server {
		Addr: ":8080",
	}
	// '/'라는 url에 대한 핸들러 함수를 등록
	// http 요청시 http://localhost:8080/ 라는 HTTP 요청 수신시, 핸들러 함수 실행으로 w과 r을 받아올 수 있다.
	// TCP 통신 주소(http)로 통신 요청을 받는다.
	// http:ListenAndServe(":port number", serveMux), serveMux가 nil일 경우 DefaultServeMux(HTTP request router) 사용
	
	if err:= server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
```

# 1. slack receiver with prometheus
- [x] (devops 채널에 대한 수신 웹훅 생성)
- [x] (웹훅 URL 준비 `https://hooks.slack.com/services/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` )
- [ ] (프로메테우스 경고 알림 룰 준비, 알림 매니저에서 슬랙 리시버 생성을 위해 prometheus-server의 configmap 수정)
- [ ] (프로메테우스 확인)
- [ ] (test-pod 배포)


#### 출처
- [간단한 웹서버 도커 이미지 만들기](https://dydtjr1128.gitbook.io/understanding-docker/2.release-docker-container/1-make-simple-docker-image)
- [Go net/http 패키지](https://jeonghwan-kim.github.io/dev/2019/02/07/go-net-http.html)
- [예제로 배우는 Go 프로그래밍](http://golang.site/go/article/111-%EA%B0%84%EB%8B%A8%ED%95%9C-%EC%9B%B9-%EC%84%9C%EB%B2%84-HTTP-%EC%84%9C%EB%B2%84)

<br/><br/><br/>
