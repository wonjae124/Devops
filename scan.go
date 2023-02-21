package main

import "fmt"

func main(){
	var a int
	var b int

	// Scan함수의 반환은 입력된 값의 개수, 입력 실패시 에러 발생
	n, err := fmt.Scanln(&a,&b)
	// 에러 발생시, nil(비어 있는 값 0X0) 대신에 에러값(0x4c6268)이 출력된다
	if err != nil{
		fmt.Println(n, err)
	}else{
		print(err)
		fmt.Println(n,a,b)
	}
}