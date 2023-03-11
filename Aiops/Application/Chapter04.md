# 0.client-go, mibikube install
- `minikube start --driver=docker`

# 1. 배포 명령어
yaml(k8s manifest file)로 리소스 생성`kubectl apply -f listner.yaml`


# 2. 코드

```go
package main

import (
	"database/sql"
	"fmt"

	"github.com/ghodss/yaml"
	_ "github.com/lib/pq"
)

func main() {
	deployment := []byte(`
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
name: my-nginx
spec:
template:
metadata:
    labels:
    run: my-nginx
spec:
    containers:
    - name: my-nginx
    image: nginx
    ports:
    - containerPort: 80`)

	// yaml to json
	jsonBytes, err := yaml.YAMLToJSON(deployment)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	jsonString := string(jsonBytes)

	//json to yaml
	yamlBytes, err := yaml.JSONToYAML(jsonBytes)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	yamlString := string(yamlBytes)

	// DB open
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "postgres", "won", "wonjae")

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	db.Exec("drop table t")
	//DB insert
	db.Exec("create table t(id serial primary key, jsonformat json, yamlformat text)")
	db.Exec("INSERT INTO t(jsonformat, yamlformat) VALUES($1, $2)", jsonString, yamlString)

	fmt.Println("Done, YAML <-> Json inserting on postgresql")
}
```
<br/><br/>

# 05. 결과물
- `Go run test.go`
	- <img src = "https://github.com/wonjae124/Devops/blob/main/image/%EC%8A%A4%ED%81%AC%EB%A6%B0%EC%83%B7%202023-03-07%2016-55-53.png?raw=true">

- dbname: wonjae의 테이블명 t 쿼리
	- `SELECT * FROM t;`
	- <img src = "https://github.com/wonjae124/Devops/blob/main/image/%EC%8A%A4%ED%81%AC%EB%A6%B0%EC%83%B7%202023-03-07%2016-59-54.png?raw=true">

<br/><br/>

# 06. 느낀점
- Go 와 sql을 단시간에 구현하며 배울 수 있어서 좋은 기회였고, 성취감이 있었다. 
- Go로 로컬의 postgresql 서버를 연동해보면서, 다른 프로그램으로의 확장성이 뛰어나다는 점을 다소 이해했다.
- GO는 Backend 분야에서 활용되는 언어임을 알게되었으며, 다른 오픈소스와 유연하게 연동이 가능하다는 점에 흥미를 느꼈고 다른 오픈소스와 또 연동을 해보고 싶어졌다

<br/><br/>

#### 출처

- [client-go library to develop Kubernetes native app](https://youtu.be/vlw1NYySbmQ)
  <br><br><br>
