# 0. kurbernetes object  
- 클러스터 접속 정보 확인 : `cat ~/.kube/config`
  - kubectl명령어를 내리면 http request를 받아줄 서버가 필요하다. 이 때, 어디로 REST API를 보낼지에 대한 정보가 config에 쓰여져 있다. 
  - 클러스터 유저 말고, 컨텍스트란? : 어떤 클러스터에 사용자가 접속할지를 조합한 정보
- 쿠버네티스 오브젝트 : manifest파일,Rest API로 전달, 쿠버네티스 클러스터를 이용해 애플리케이션을 배포하고 운영하기 위해 필요한 모든 쿠버네티스 리소스
  -   어떤 애플리케이션(pod)
  -   얼마나 (replicaset)
  -   어디에(node, namepsace)
  -   어떤 방식으로 배포(deployment)
  -   트래픽을 어떻게 호출하고 로드밸런싱(service, endpoints)
 -  쿠버네티스 클러스터에 사용 가능한 오브젝트 목록 조회 : `kubectl api-resources`
 -  쿠버네티스 오브젝트 설명, 1레벨 속성 설명 `kubectl explain <type>`
 -  쿠버네티스 오브젝트 속성의 구체적인 설명, json 경로 :`kubectl explain <type>.<filedName>[.<filedName>] 

- pods : 컨테이너
  - 여러 컨테이너를 감싸고 있는 콩껍질을 팟이라고 한다.
  - 노드에서 컨테이너를 실행하기 위한 가장 기본적인 배포 단위
  - 여러 노드에 1개 이상의 pod을 분산 배포/실행 가능(Pod replicas set) 
  - 하나의 노드에 여러 개의 팟이 존재할 수 있다  
  - 쿠버네티스는 pod을 생성할 때, 노드에서 유일한 IP 할당(서버 분리, 팟은 고유한 하나의 서버다)
    - Pod 내부 컨테이너 간에 localhost로 통신 가능, 포트 충돌 주의(다른 팟끼리 통신하려면 pod IP를 알아야한다)
    - pod 안에서 네트워크와 볼륨 등 자원을 공유
  - podIP는 클러스터 안에서만 접근 가능함. -> 만약, 클러스터 외부 트래픽을 받기 위해서는 service 혹은 ingress 오브젝트 필요
    - 클러스터 외부에서 원래는 접근 가능, 내부에서 접근하기 위해서는 별도의 오브젝트 필요
    - Pod IP는 컨테이너와 공유된다. 따라서, 컨테이너 간 포트 충돌 주의
    - 하나의 pod에 속한 컨테이너들은 localhost로 통신할 수 있다

  - Pod: container = 1:1 or 1:N 결정
    - 1. 컨테이너들의 라이프사이클이 서로 같은가?
    - 2. 스케일링 요구사항이 같은가? - 웹 서버 vs 데이터베이스, 트래픽이 많은가 vs 그렇지 않은가
    - 3. 인프라 활용도가 더 높아지는 방향으로 설계(쿠버네티스가 노드 리소스 등 여러 상태를 고려하여 pod을 스케쥴링)
  - pod에 환경변수를 담아 컨테이너에 전달할 수 있다
  
- service 오브젝트
  - 클러스터 외부에서 접근할 수 있는 고정적인 단일 엔드포인트 필요
  - Pod 집합을 클러스터 외부로 노출하기 위한 service 오브젝트 도입
 
- Replicaset 오브젝트
  - 나도 모르는 사이에 node가 꺼질 수 있음. pod은 자가 치유 능력이 없다
  - 사용자가 선언한 수만큼 pod을 유지해주는 Replicaset 오브젝트 도입
  
  
  
# 1. option
- pod ip 확인 : `kubectl get pod -o wide`
- json형식으로 pod 확인 : `kubectl get pod hello-app -o json`
- 컨테이너 내부 들여다보기 : `kubectl exec hello-app --cat /etc/json
- 환경변수 확인 : `kubectl exec hello-app -- env`
- 컨테이너의 리스닝하고 있는 pod 확인 :`kubectl exec hello-app -- netstat -an
- 로컬포트의 8080과 파드의 8080을 트래픽 연결하기 : kubectl port-forward hello-app 8080:8080
- 리퀘스트, 리스폰드 정보 확인법 : curl -v locatlhost:8080 
- pod 전체 삭제 : kubectl delete pod --all
- 클러스터 세팅 확인 : `kubectl config current-context`
- 파드 내의 컨테이너 로그 확인 : `kubectl log blue-green-app -c blue-app`
- 컨테이너의 환경변수 확인 : `kubectl exec blue-green-app -c blue-app -- printenv POD_IP NAMESPACE NODE_NAME`
- blue-app 컨테이너 -> green-app 컨테이너 /tree, /hello 요청 실행
- kubectl exec blue-green-app -c blue-app -- curl -vs localhost:8081/tree
- kubectl exec blue-green-app -c blue-app -- curl -vs localhost:8081/hello

- green-app 컨테이너 -> blue-app 컨테이너 /sky, /hello 요청 실행
  - `kubectl exec blue-green-app -c green-app -- curl -vs localhost:8080/sky`
  - `kubectl exec blue-green-app -c green-app -- curl -vs localhost:8080/hello`

- blue-app 컨테이너 -> red-app 컨테이너(포트번호 8080임). 앱 /rose, /hello 요청 실행
  - `export RED_POD_IP=$(kubectl get pod red-app -o jsonpath="{.status.podIP}")`
  - `echo $RED_POD_IP`
  - `kubectl exec blue-green-app -c blue-app -- curl -vs $RED_POD_IP:8080/rose`
  - `kubectl exec blue-green-app -c blue-app -- curl -vs $RED_POD_IP:8080/hello`

- red-app 컨테이너 -> blue-app 컨테이너 /sky, /hello 요청 실행
  - `export BLUE_POD_IP=$(kubectl get pod blue-green-app -o jsonpath="{.status.podIP}")`
  - `echo $BLUE_POD_IP`
  - `kubectl exec red-app -- curl -vs $BLUE_POD_IP:8080/sky`
  - `kubectl exec red-app -- curl -vs $BLUE_POD_IP:8080/hello`

- 포트포워딩을 통해 웹브라우저로 각 컨테이너 요청/응답 확인
- 로컬호스트의 8080을 컨테이너의 8080으로 넘긴다
  - `kubectl port-forward blue-green-app 8080:8080`
  - `kubectl port-forward blue-green-app 8081:8081`
  - `kubectl port-forward red-app 8082:8080`



# 2. 코드

```
# Pod API 버전: v1
# Pod 이름: hello-app
# Pod 네임스페이스: default
# 컨테이너 이름/포트: hello-app(8080)
# 도커 이미지: yoonjeong/hello-app:1.0
# 환경변수:
# -- POD_NAME(metadata.name), POD_IP(status.podIP)
# -- NAMESPACE_NAME(metadata.namespace)
# -- NODE_NAME(spec.nodeName), NODE_IP(status.hostIP)
# -- STUDENT_NAME(본인이름), GREETING(STUDENT_NAME을 참조한 인삿말)
apiVersion: v1
kind: Pod
metadata:
  name: hello-app
  namespace: default
spec:
  containers:
    - name: hello-app
      image: yoonjeong/hello-app:1.0
      ports:
        - containerPort: 8080
      env:
        - name: POD_NAME
          valueFrom:
            fieldRef:
              fieldPath: metadata.name
        - name: POD_IP
          valueFrom:
            fieldRef:
              fieldPath: status.podIP
        - name: NAMESPACE_NAME
          valueFrom:
            fieldRef:
              fieldPath: metadata.namespace
        - name: NODE_NAME
          valueFrom:
            fieldRef:
              fieldPath: spec.nodeName
        - name: NODE_IP
          valueFrom:
            fieldRef:
              fieldPath: status.hostIP
        - name: STUDENT_NAME
          value: 정원재
        - name: GREETING
          value: 안녕하세요. $(STUDENT_NAME)
      resources:
        limits:
          memory: "128Mi"
          cpu: "500m"


```
<br/><br/>

# 05. 결과물


#### 출처
- [Fast campus - DevOps 마스터 Kit
with Linux, Kubernetes, Docker](https://fastcampus.co.kr/dev_online_awsdevops/?utm_source=google&utm_medium=cpc&utm_campaign=hq%5E210101%5E206717&utm_content=devops&utm_term=&gclid=Cj0KCQjwk7ugBhDIARIsAGuvgPYd_OMj-l6E9FlFSCgKHbwiiCpfnQrQUp9_o8FuWSjS8tkRuZwrCJ4aAqxbEALw_wcB)
<br><br><br>
