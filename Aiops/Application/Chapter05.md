# 0.Job with client go
- 클라이언트를 생성해서, K8s api server에게 talk를 한다.


# 1.  
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
  - Pod: container = 1:1 or 1:N 결정
    - 1. 컨테이너들의 라이프사이클이 서로 같은가?
    - 2. 스케일링 요구사항이 같은가? - 웹 서버 vs 데이터베이스, 트래픽이 많은가 vs 그렇지 않은가
    - 3. 인프라 활용도가 더 높아지는 방향으로 설계(쿠버네티스가 노드 리소스 등 여러 상태를 고려하여 pod을 스케쥴링)
  
# 2. 

# 3. 코드

```golang

```
<br/><br/>

# 05. 결과물


#### 출처
- [create-kubernetes-jobs-in-golang](https://youtu.be/vlw1NYySbmQ](https://dev.to/narasimha1997/create-kubernetes-jobs-in-golang-using-k8s-client-go-api-59ej)
<br><br><br>
