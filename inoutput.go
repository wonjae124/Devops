package main

import "fmt"

func main(){
	var a int = 10
	var b int = 20
	var f float64 = 32799438743.8297
	
	// 개행 X
	fmt.Print("a:",a,"b:",b)
	// 출력값 사이 공란, 개행 O
	fmt.Println("a:",a,"b:",b,"f:",f)
	// 출력값 사이 공란, 개행 X, 서식(길이) 지정
	fmt.Printf("a: %d b: %d f: %f\n",a,b,f)
}
