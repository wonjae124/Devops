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
- 
# 2. 삭제 명령어
- yaml deployment 삭제
	- `kubectl delete -f lister.yaml`
	-`kubectl delete deployment lister`
- 테이블 삭제 : `DROP TABLE p;`

# 3. 코드

```golang
// 순서
// 1. minkube start
// 2. dockerfile로 image build 하여, 이미지 생성함. 이미지 이름은 lister:0.1.0
// 3. dockerpush로 이미지를 로그인된 도커 계정의 허브에 올리기
// 4. kubectl create deployment 이미지 이름으로 lister.yaml 파일 생성
// 5. kubectl create -f lister.yaml 으로 pod 배포
// 6. rest.InclusterConfig()를 통해, 내 프로그램을 k8s에 접속시켜서 config를 받아온다.
// 7. 팟, 디플로이먼트 리스트를 불러오게끔, role을 지정
// 8. default namspace의 default service account에 role binding 지정
// 9. postgresql과 연동해서 입력하도록 코드 추가

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
	// k8s 어플리케이션을 host machine에서 돌릴 때는, ./kube/config가 필요하지만, 그게 아니라 클러스터 내부에서 통신은 config를 필요로 하지 않는다. rest 패키지를 쓴다. 
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

# 06. 느낀점
- Go 와 sql을 단시간에 구현하며 배울 수 있어서 좋은 기회였고, 성취감이 있었다. 
- Go로 로컬의 postgresql 서버를 연동해보면서, 다른 프로그램으로의 확장성이 뛰어나다는 점을 다소 이해했다.
- GO는 Backend 분야에서 활용되는 언어임을 알게되었으며, 다른 오픈소스와 유연하게 연동이 가능하다는 점에 흥미를 느꼈고 다른 오픈소스와 또 연동을 해보고 싶어졌다
<br/><br/>

#### 출처
- [youtube - client-go library to develop Kubernetes native app](https://youtu.be/vlw1NYySbmQ)
- [kubuctl cheet sheet](https://kubernetes.io/docs/reference/kubectl/cheatsheet/)
- [youtube - Run your client-go application as a pod in cluster](https://youtu.be/NeV-jR_LssA)
  <br><br><br>
