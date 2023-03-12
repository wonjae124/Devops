# 0.client-go 
- 클라이언트를 생성해서, K8s api server에게 talk를 한다.


# 1. 명령어
- 
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
	"k8s.io/client-go/tools/clientcmd" //  
)
func main() {
	// k8s 어플리케이션과 외부 원격 클러스터의 연결은 ./kube/config가 필요하다. 만약, 클러스터 내부의 통신은 config를 필요로 하지 않는다. rest 패키지를 쓴다. 
	
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


#### 출처
- [create-kubernetes-jobs-in-golang](https://youtu.be/vlw1NYySbmQ](https://dev.to/narasimha1997/create-kubernetes-jobs-in-golang-using-k8s-client-go-api-59ej)
<br><br><br>
