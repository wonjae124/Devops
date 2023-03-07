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
	
	fmt.Println("Done, YAML to JSON, JSON to YAML with inserting postgres")
}
