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

- 한계점으로, yaml와 json을 변환하려면, yaml과 json 모두 사전에 struct 정의가 필요하기에, 매번 하드코딩하는건 부적합하다고 판단함
- 만약, yaml이 배포용도의 파일이면, 띄어쓰기와 들여쓰기가 중요함. 하지만, yaml을 json으로 변환할시 indent가 사라지는 문제 존재. 이에,  운영 환경에서, 호환 불가 문제 존재할 것으로 예상
- 이에, yaml과 json 변환 패키지 탐색 필요성을 느낌
<br/><br/>

### 출처

- [github - ghodss/yaml](https://github.com/ghodss/yaml)
- [Naver blog - Golang Marshal, Unmarshal 차이](https://etloveguitar.tistory.com/44)
- [Tistory blog - Marshal, Unmarshal 함수](https://codecollector.tistory.com/1513)
<br><br><br>
