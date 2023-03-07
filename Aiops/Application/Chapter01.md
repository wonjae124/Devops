
# 0.Yaml, Json on Golang

## Yaml vs json
- Json은 java machine object로 구성되어, 확장성이 좋다. 외부로 공개 가능함. 
- Yaml은 시스템 구성용도로 쓰이며 외부로 공개하지 않음
- Json은 serilization format 이다
- Yaml과 Json은 공통적으로 key:value로 구성된다 또한 json parsing을 위해서, 별도로 yaml parsing 후에 yaml의 하위 집합으로 표현한다.
- 서로 다른 점으로 Json은 Key를 double quotes로 반드시 감싸야한다. yaml은 그럴 필요가 없다.
- golang에서 사용하기 위해서, marshal, unmarshal로 변환이 필요하다.

# 1. Marshal, UnMarshal
- Marshal   : Go Object(struct,string) => []byte, string
- Unmarshal : []byte, string => Go object(struct,string)

- marshal과 unmarshal 과정을 이해하고자 간단한 예제 연습
  - String(Json string) -> Go object(custom Struct)
```
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
	jsonString := `{"name": "battery sensor","capacity":40, "time":
	"2019-01-21T19:07:28Z", "info": {
		"desc":"a sensor reading"
	}}`

	var reading SeonsorReading
	err := json.Unmarshal([]byte(jsonString),&reading) 
	if err != nil {
		fmt.Println(err)
	} 

	fmt.Printf("%+v\n",reading)
}
```


### 출처

- [Linuxhint - Yaml vs json](If you want to parse JSON then you have to use the YAML parser as JSON is a subset of YAML)
- [Naver blog - Golang Marshal, Unmarshal 차이](https://etloveguitar.tistory.com/44)
<br><br><br>
