package main

import "fmt"

func main(){
	// 1. variable 변수명 자료형(타입) = 초깃값
	// 2. 상수로만 초기화 하는 변수 = 정적 변수(static variable)
	var a int = 10 
	var b int = 30
	
	// 3. 타입 생략, 변수 타입은 우변의 타입
	var c = 4
	
	// 4. 변수 선언문은 변수 선언 키워드와 생략해도 자동으로 만듦
	e :=5

	var msg string = "Hello Variable"

	a = 20
	// 5. 수식 선언과 함께, 연산자 수식결과로 초기화
	var d = a*b 
	msg = "Good Morning"
	fmt.Println(msg,a,c,d,e)
}

/* 

변수명 규칙

1. 첫글자는 문자 or _ 사용함. 숫자 불가능
2. _를 제외한 특수문자(space 포함) 불가
3. 다른 언어 문자(한글) 지양
4. 변수명에 여러 단어가 이어지면, 두 번째 단어부터는 대문자 권장
5. 변수명은 되도록 짧게, 임시 로컬 변수는 한 글자 권장
6. 밑줄은 일반적으로 변수명에 사용 X, _를 사용하는 경우는 함수와 패키지 부분에서 설명


*/