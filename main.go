package main

import (
	"fmt"
	"time"
)

func readword(ch chan string) {
	fmt.Println("Type a word, then hit Enter.")
	var word string
	fmt.Scanf("%s", &word) // 입력을 받는다
	ch <- word             // 입력을 채널로 전달
}

func printchar() { // 2초에 한 번씩 점을 찍는다
	for { // for는 while문으로 보면 된다
		fmt.Printf(".")
		time.Sleep(2 * time.Second)
	}
}

func main() {
	defer fmt.Println("===== BYE..") // defer는 코드를 지연해서 실행, 앱이 종료되면 마지막에 실행된다.
	go printchar()                   // 함수명 앞에 go 써주면 동시에 실행된다.

	ch := make(chan string) // make - go를 이용해 쓰레드를 만들고, 쓰레드 간에 통신(메세지)하는 채널을 만든다
	go readword(ch)

	select { // switch역할, 채널을 받는다
	case word := <-ch:
		fmt.Println("\nReceived: ", word) // 메세지를 받아, 바로 출력
	}
}
