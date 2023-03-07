# 0.Handle postgresql on Golang

- 목표 : postgresql, golang 연동하여 yaml<->json 입력

# 1. sql package

- database/sql은 여러 종류의 sql 데이터베이스 지원함. 드라이버 사용 가능
- sql.open, 객체 DB생성
- defer db.Close, 시간 지연 후 DB 닫기
- db.Query, 쿼리 선택
- db.Exec, insert 실행
  <br/><br/>

# 2. postgres

- 스키마

  - 개체들의 논리적 집합, table, view, function, index, data type. operator<br/>
  - 스키마는 각각의 이름으로 식별할 수 있음<br/>
  - 데이터생성시 기본 스키마는 public임

- user : default
- dbname : wonjae
- postgres control option
  - `\dn+ : 스키마 권한`<br/>
  - -`\l : list of database`<br/>
  - -`\d : list of relations`<br/>
  - -`\dt : list of relations(type : table)  => 테이블 조회`<br/>
  - `\connect DBNAME : switch other database`<br/>
    <br/><br/>

# 3. Panic

- postgresql relation does not exist
  - 원인 : access privilege, schema 생성을 안함
  - $ find / -name pg_hba.conf
  - $ sudo vim /etc/postgresql/14/main/pg_hba.conf
  - peer, md5 ⇒ trust로 변경
- schema "wonjae" does not exist
  - 원인 : dbname이 postgres이기에 schema가 없다는 에러 발생
  - 해결 : wonjae라는 dbname으로 schema 생성
    <br/><br/>

- db.create는 자동으로 schema가 생성되지 않으므로 별도 스키마 생성 필요하다는걸 느꼈다

# 4. 코드

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

- [Hevodata - How to use Golang PostgreSQL?](https://hevodata.com/learn/golang-postgres/)
- [예재로 배우는 Go 프로그래밍 -SQL DB 활용](http://golang.site/go/article/106-SQL-DB-%ED%99%9C%EC%9A%A9)
- [GIS DEVELOPER blog - [Golang] PostgreSQL 다루기](http://www.gisdeveloper.co.kr/?p=2456)
  <br><br><br>
