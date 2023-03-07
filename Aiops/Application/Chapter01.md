
# 0.Yaml, Json on Golang

## Yaml vs json
- Json은 java machine object로 구성되어, 확장성이 좋다. 외부로 공개 가능함. 
- Yaml은 시스템 구성용도로 쓰이며 외부로 공개하지 않음
- Json은 serilization format 이다
- Yaml과 Json은 공통적으로 key:value로 구성된다 또한 json parsing을 위해서, 별도로 yaml parsing 후에 yaml의 하위 집합으로 표현한다.
- 서로 다른 점으로 Json은 Key를 double quotes로 반드시 감싸야한다. yaml은 그럴 필요가 없다.
- golang에서 사용하기 위해서, marshal, unmarshal로 변환이 필요하다.<br/><br/>

# 1. Marshal, UnMarshal
- Marshal   : Go Object(struct,string) => []byte, string
- Unmarshal : []byte, string => Go object(struct,string)

- marshal과 unmarshal 과정을 이해하고자 간단한 예제 연습
  - String(Json string) -> Go object(custom Struct)
  - `func Marshal(v interface{}) ([]byte, error),      입력 = 정수 또는 구조체, 출력 = 바이트 리스트`<br/>
  - `func Unmarshal(data []byte, v interface{}) error, 입력 = 바이트 리스트, 출력 = 정수 또는 구조체`


```go


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

type test struct{
	Stmt string `json:"stmt"`
	Num int `json:"age"`
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

	fmt.Printf("%+v\n",reading) // {Name:battery sensor Capacity:40 Time:2019-01-21T19:07:28Z Information:{Description:a sensor reading}}
	
	/* 
	struct 형태인걸 다시 string으로 변환 불가. 아래는 error
	fmt.Printf("%+v\n",string(reading))
	=> cannot convert reading (variable of type SeonsorReading) to type string
	*/
	
	var i = test{"WHY",1}
	bytes, _ := json.Marshal(i);
	fmt.Println(bytes) // [123 34 115 116 109 116 34 58 34 87 72 89 34 44 34 97 103 101 34 58 49 125]
	fmt.Println(string(bytes)) // {"stmt":"WHY","age":1}
	
}
```
<br/><br/>
# 2. yaml to json, json to yaml on Golang
<img src = "https://github.com/wonjae124/Devops/blob/main/image/%EC%8A%A4%ED%81%AC%EB%A6%B0%EC%83%B7%202023-03-07%2015-59-11.png?raw=true" width=1000>

```go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	_ "github.com/lib/pq"
	"gopkg.in/yaml.v3"
)

type operatingSystems struct{
	Windows10			int `yaml:"Windows 10"`
	WindowsServer2019 	int `yaml:"Windows Server 2019"`
	WindowsServer2022  	int `yaml:"Windows Server 2022"`
	MacOS 				int `yaml:"MacOS"`
	CentOS 				int `yaml:"CentOS"`
	Photon				int `yaml:"Photon"`
}

type operatingSystemsjson struct {
	Windows10 			int `json:"Windows 10"`
	WindowsServer2019 	int `json:"Windows Server 2019"`
	WindowsServer2022  	int `json:"Windows Server 2022"`
	MacOS 				int `json:"MacOS"`
	CentOS 				int `json:"CentOS"`
	Photon				int `json:"Photon"`
} 

func main(){
	fmt.Println("Parsing YAML file")
	var yamlfileName string = "/home/won/temp/origin-yaml.yml"


	// 1. Yaml to Json
	// pwd : /home/won/바탕화면/go_test/sample-app
	// yaml file : /home/won/temp/operating-systems.yml

	// Yaml file -> yaml byte
	yamlFile, err := ioutil.ReadFile(yamlfileName)
		if err != nil{ 
			fmt.Printf("Error reading YAML file : %s\n",err)
			return
		}
		// initial yaml struct variables 
		var oses operatingSystems 
		// Unmarshal, yaml byte -> yaml struct variables 
		yaml.Unmarshal(yamlFile, &oses) 


	// yaml struct variables -> json struct variabels 

	var osesjson = operatingSystemsjson{
		Windows10: 			oses.Windows10,
		WindowsServer2019:	oses.WindowsServer2019,
		WindowsServer2022: 	oses.WindowsServer2022,
		MacOS: 				oses.MacOS,
		CentOS:				oses.CentOS,
		Photon:				oses.Photon,
		}
		
	// marshal, json struct variables -> json byte
	jsonOutput, err := json.Marshal(osesjson)
	// sqlInfo := fmt.Sprintf("Result: %+v\n",osesjson) 
		fmt.Printf("Result1: %+v\n",osesjson) 

		// json byte -> file
		// 0644 소유자는 읽기,쓰기 권한, 그룹/기타 사용자는 읽기 권한
	err = ioutil.WriteFile("./Go-operating-systems.json",jsonOutput,os.FileMode(0644))

	// 2. Json to Yaml
	// pwd : /home/won/바탕화면/go_test/sample-app
	// json file : /home/won/temp/operating-systems.json

	fmt.Println("Parsing JSON file")
	// json byte 
	var jsonFilename string = "/home/won/temp/origin-json.json"
	
	jsonFile, err := ioutil.ReadFile(jsonFilename)
		if err != nil {
			fmt.Printf("Error reading JSON FILE : %s\n",err)
			return
		}
		// initial json struct variable
		var pses operatingSystemsjson
		// json byte -> json struct variable
		json.Unmarshal(jsonFile,&pses)

	// json struct variable -> yaml struct variables
	var psesyaml = operatingSystems{
		Windows10: 		    pses.Windows10,
		WindowsServer2019:  pses.WindowsServer2019,
		WindowsServer2022:  pses.WindowsServer2022,
		MacOS: 				pses.MacOS,
		CentOS:				pses.CentOS,
		Photon:				pses.Photon,	
	}

	// yaml struct variables -> yaml bytes
	yamlOutput, err := yaml.Marshal(psesyaml)
		fmt.Printf("Result2:%+v\n", psesyaml)

	err = ioutil.WriteFile("./Go-operating-systems.yaml",yamlOutput,os.FileMode(0644))
	fmt.Printf("please:%+v\n", string(yamlOutput))
}	
```
<br/><br/>
# 03. 느낀점

- 한계점으로, yaml와 json을 변환하려면, yaml과 json 모두 사전에 struct 정의가 필요하기에, 매번 하드코딩하는건 부적합하다고 판단함
- 만약, yaml이 배포용도의 파일이면, 띄어쓰기와 들여쓰기가 중요함. 하지만, yaml을 json으로 변환할시 indent가 사라지는 문제 존재. 이에,  운영 환경에서, 호환 불가 문제 존재할 것으로 예상
- 이에, yaml과 json 변환 패키지 탐색 필요성을 느낌
<br/><br/>

### 출처

- [Linuxhint - Yaml vs json](https://linuxhint.com/yaml-vs-json-which-is-better/)
- [Naver blog - Golang Marshal, Unmarshal 차이](https://etloveguitar.tistory.com/44)
- [Tistory blog - Marshal, Unmarshal 함수](https://codecollector.tistory.com/1513)
<br><br><br>
