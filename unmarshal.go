package main

import (
	"encoding/json"
	"fmt"
)

type SeonsorReading struct{
	Name string `json:"name"`
	Capacity int `json:"capacity"`
	Time string `json:"time"`
	Information Info `json:"info"`
}

type Info struct {
	Description string `json:"desc"`
}

func main(){
	fmt.Println("Hello world") 
	jsonString := `{"name": "battery sensor","capacity":40, "time":
	"2019-01-21T19:07:28Z", "info": {
		"desc":"a sensor reading"
	}}`

	var reading SeonsorReading
	err := json.Unmarshal([]byte(jsonString),&reading) // 이거 배열 만드는거?, 뒤에 & 연산자 왜 붙지? 주소?
	if err != nil {
		fmt.Println(err)
	} // nil이 뭐지

	fmt.Printf("%+v\n",reading)
}