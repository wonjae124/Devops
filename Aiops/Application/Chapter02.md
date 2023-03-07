# 0.Yaml to Json, Json to Yaml on Golang
- json은 yaml로 표현 할 수 있음
- 기존 go-yaml 패키지는 yaml의 유의미한 비어있는 값에 null로 값을 생성하는 문제 존재
- 해당 문제를 개선 및 사용하기 간편한 패키지로, ghodss/yaml 존재하므로 선택함
<br/><br/>
# 1. ghodss/yaml 
- github 
  - star 980
  - 이슈 29
  - 최근 업데이트 June, 2022
  <br/>
- ghodss/yaml 패키지 이해를 위해, 간단한 예제 실습
  - `func JSONToYAML(j []byte) ([]byte, error)`
  - `func YAMLToJSON(y []byte) ([]byte, error)`
  - 입력 : 바이트 리스트
  - 출력 : 바이트 리스트
<br/>  <br/>
```go

package main

import (
	"fmt"
	"github.com/ghodss/yaml"
)

func main() {
	j := []byte(`{"name": "John", "age": 30}`)
	y, err := yaml.JSONToYAML(j)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Println(string(y))
	/* Output:
	name: John
	age: 30
	*/
	j2, err := yaml.YAMLToJSON(y)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Println(string(j2))
	/* Output:
	{"age":30,"name":"John"}
	*/
}
}
```
<br/><br/>

# 02. 느낀점
- 새로운 패키지를 사용할 때는, 함수의 입력타입과 출력타입을 조사해두어야겠다. 이걸 잘 몰라서, postgres에 어떻게 입력할지에 대해서 다소 헤맸다  
<br/><br/>

### 출처

- [github - ghodss/yaml](https://github.com/ghodss/yaml)
<br><br><br>
