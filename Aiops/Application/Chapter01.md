
# 0.Golang yaml to json

## Yaml vs json
### Json은 확장성이 좋으며, 외부로 공개 가능함. 
### Yaml은 시스템 구성용도로 쓰이며 외부로 공개하지 않음
### Json은 serilization format으로, marshal, unmarshal로 이루어진다.
### Yaml과 Json은 공통적으로 key:value로 구성되며, 기본적으로 json은 yaml으로 parsing해서 표현이 된다
### 서로 다른 점으로 Json은 Key를 double quotes로 반드시 감싸야한다. yaml은 그럴 필요가 없다.

- postgres account 실행
- user로 sammy 추가<br/><br/><br/>
```
$ sudo apt install postgresql postgresql-contrib 

$ sudo systemctl start postgresql.service

$ go mod init tutorial 

$ sudo -u postgres createuser --interactive

$ sudo -u postgres psql 

$ sudo adduser wonjae

$ sudo -i -u wonjae psql

```


### 출처

- [Linuxhint - Yaml vs json](If you want to parse JSON then you have to use the YAML parser as JSON is a subset of YAML)<br><br><br>
