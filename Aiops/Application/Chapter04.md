# 0.client-go, mibikube install
- `minikube start --driver=docker`

# 1. 명령어
- yaml(k8s manifest file)로 pod lister를 배포`kubectl apply -f lister.yaml`
- main.go, go.mod, go.sum을 통합해서 실행파일을 만들기 `go build`
- pod과의 통신(interacting) `kubectl logs $(pods name)` ex) kubectl logs lister-69685658b8-jrlsd
- kubectl api role, rolebinding 확인 `kubectl api-resources`
- 롤 만들기 : `kubectl create role poddepl --resource pods,deployments --verb list`
- 롤 바인딩 : `kubectl create rolebinding poddepl --role poddepl --serviceaccount default:default`
- 이미지 만들기 :`docker build -t $(image name)`
- 이미지 이름 변경 : `docker tag ~`
- 이미지 도커 허브로 보내기 : `docker push ~ `
- 지속적인 노드, 파드, 디플리먼트  : watch kubectl get nodes or podes or deployments

# 2. 삭제 명령어
- yaml deployment 삭제
	- `kubectl delete -f lister.yaml`
	-`kubectl delete deployment lister`
- 테이블 삭제 : `DROP TABLE p;`

# 3. 코드

```golang

package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"time"

	_ "github.com/lib/pq"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes" // k8s.io/api v0.26.1, k8s apimachinary 설치
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd" //  기본적인 golang json, yaml, log 설치
)
func main() {
	// k8s 어플리케이션과 외부 원격 클러스터의 연결은 ./kube/config가 필요하다. 그게 아니라 클러스터 내부에서 통신은 config를 필요로 하지 않는다. rest 패키지를 쓴다. 
	kubeconfig := flag.String("kubeconfig","/home/won/.kube/config","location to your kubeconfig file")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	config.Timeout = 120 * time.Second
 	if err != nil{
		fmt.Printf("error %s building config from flags\n",err.Error())
		config, err = rest.InClusterConfig()
		if err != nil {
			fmt.Printf("error %s, geting inclusterconfig", err.Error())
		}
	}
		// postgresql
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
							"localhost",5432,"postgres","won","wonjae")
				
	db, err := sql.Open("postgres",psqlInfo)
	db.Exec("drop table p")

	if err != nil {
		panic(err)
	}
	db.Exec("CREATE TABLE p(id serial primary key, pod text, deployment text)")

	//runtime.Object

	clientset, err := kubernetes.NewForConfig(config) // 새로운 포드 리스트 확인해서 삭제, 업데이트, 디플레이먼트하게끔 인터렉트 가능
	if err != nil {
		fmt.Printf("error %s creating clientset\n",err.Error())
	}
	ctx := context.Background()
	pods, err := clientset.CoreV1().Pods("default").List(ctx,metav1.ListOptions{}) // Pods("네임스페이스 이름")
	if err != nil {
		fmt.Printf("error %s while listing all the pods form default namespace\n",err.Error())
	}
	// pod resource와 api를 통해 디플로이먼트 가능
	// namespace로 default 선택 
	fmt.Println("Podes from default namepsace")
	for _,pod := range pods.Items{
		podData := fmt.Sprintf("%s", pod.Name)
		db.Exec("INSERT INTO p(pod) VALUES($1)", podData)
	}

	fmt.Println("Deployments are ")
	deployments, err := clientset.AppsV1().Deployments("default").List(ctx,metav1.ListOptions{}) // Deployments("네임스페이스 이름")
	if err != nil{
		fmt.Printf("listing deployments %s \n",err.Error())
	}
	
	for _, deployment := range deployments.Items{
		deploymentData := fmt.Sprintf("%s",deployment.Name)
		db.Exec("INSERT INTO p(deployment) VALUES($1)", deploymentData)
	}

	db.Close()
	// db.Exec("INSERT INTO t(pod, deployment) VALUES($1, $2)", podData, deploymentData)
	fmt.Println("Done")

}
```
<br/><br/>

# 05. 결과물
- <img src = "https://github.com/wonjae124/Devops/blob/main/image/%EC%8A%A4%ED%81%AC%EB%A6%B0%EC%83%B7%202023-03-11%2017-54-52.png">
<br/><br/>


#### 출처
- [youtube - client-go library to develop Kubernetes native app](https://youtu.be/vlw1NYySbmQ)
- [kubuctl cheet sheet](https://kubernetes.io/docs/reference/kubectl/cheatsheet/)
- [youtube - Run your client-go application as a pod in cluster](https://youtu.be/NeV-jR_LssA)
- [How to list Kubernetes Pods using Golang SDK - Faizan Bashir](https://faizanbashir.me/how-to-list-kubernetes-pods-using-golang-sdk)
<br><br><br>
