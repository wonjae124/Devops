# 0. metric 
kubelet으로부터 리소스 매트릭을 수집
이를, k8s api 서버에 Metrics Api 형태로 노출된다.
Metrics api는 자원의 사용량에 따라, pod의 개수 및 용량을 제한하는 오토 스케일링을 기능을 한다.
<br/><br/>
# 1. Kubenetes object 'Service'
- 운영환경에서 Pod 한계점
  - Pod IP는 종료, 생성시 매번 바뀐다. 클라이언트 입장에서는 Pod의 Ip를 계속 관리해야한다.
  - 이미지, 레이블 변경하여 파드를 배포한 경우에는 IP가 바뀜
  - 클라이언트는 변경된 IP를 제공받아야 한다.
  - 이를 위해 우리는 일정한 endpoint를 제공해서, pod과 통신하게 만든다.
  - 
- Service 개념과 특징, 활용 방법
  - Pod을 외부로 노출하는 다양한 Service 타입
  - 기존 Pod IP는 클러스터 내부에서만 접근 가능, 이에 클러스터 외부에서 접근할 수 있는 방법 필요.
  - 파드 추상화 = 파드들의 단일 엔드포인트, 로드벨런싱 제공
  - 파드 클라이언트는 Service IP:Port를 이용해서 파드와 통신 가능
  - Service는 Selector에 의해 선택된 파드 집합 중 임의의 파드로 트래픽을 전달한다.
      <img src="https://github.com/wonjae124/Devops/blob/main/image/%EC%8A%A4%ED%81%AC%EB%A6%B0%EC%83%B7%202023-03-15%2019-41-47.png" width=800>
 
- selector를 통해서 mapping이 된다? 이를 확인해보려면, kubectl get endpoints -n snackbar가 필요하다
- spec이 뭐더라 
- yaml 파일 작성시 띄어쓰기 두번으로 코드 구분해야한다... Tab 썼다가 실수한다.
<br/><br/>
# 2. 코드

```
# snackbar 네임스페이스 생성
kubect create namespace snackbar

# Service, Deployment 배포 (service.yaml)
kubectl apply -f service.yaml

# snackbar 네임스페이스의 모든 리소스 조회
kubectl get all -n snackbar

# snackbar 네임스페이스의 order, payment Service 상세 확인
kubectl get svc order -o wide -n snackbar
kubectl get svc payment -o wide -n snackbar

# snackbar 네임스페이스의 모든 Endpoints 리소스 확인 (kubectl get pod -o wide 조회 후 IP 비교)
kubectl get endpoints -n snackbar

# order Service IP(ClusterIP) 조회 (-o jsonpath="{.spec.clusterIP}")
kubectl get svc order -o jsonpath="{.spec.clusterIP}" -n snackbar

# 로컬에서 Service IP와 Port 호출 확인
curl $(kubectl get svc order -o jsonpath="{.spec.clusterIP}" -n snackbar)
```
<br/><br/>
# 3. service.yaml
```
apiVersion: v1
kind: Service
metadata:
  name: order
  namespace: snackbar
  labels:
    service: order
    project: snackbar
spec:
  type: ClusterIP
  selector:
    service: order
    project: snackbar
  ports:
    - port: 80
      targetPort: 8080

---
apiVersion: v1
kind: Service
metadata:
  name: payment
  namespace: snackbar
  labels:
    service: payment
    project: snackbar
spec:
  type: ClusterIP
  selector:
    service: payment
    project: snackbar
  ports:
    - port: 80
      targetPort: 8080

---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: order
  namespace: snackbar
  labels:
    service: order
    project: snackbar
spec:
  replicas: 2
  selector:
    matchLabels:
      service: order
      project: snackbar
  template:
    metadata:
      labels:
        service: order
        project: snackbar
    spec:
      containers:
        - name: order
          image: yoonjeong/order:1.0
          ports:
            - containerPort: 8080
          resources:
            limits:
              memory: "64Mi"
              cpu: "50m"

---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: payment
  namespace: snackbar
  labels:
    service: payment
    project: snackbar
spec:
  replicas: 2
  selector:
    matchLabels:
      service: payment
      project: snackbar
  template:
    metadata:
      labels:
        service: payment
        project: snackbar
    spec:
      containers:
        - name: payment
          image: yoonjeong/payment:1.0
          ports:
            - containerPort: 8080
          resources:
            limits:
              memory: "64Mi"
              cpu: "50m"

```
<br/><br/>
# 4. 환경변수를 이용한 Service 호출 
 
- 환경변수에 접근하려면, 컨테이너에 네부에 들어가야한다.
- pod 컨테이너로 명령어 전달방법
  - `kubectl exec <pod name> -n <namespace> -- <command>`
- pod 컨테이너 환경변수 확인
  - `kubectl exec <pod-name> -n <namespace> -- env | grep <pattern>` 
  - 
```
# snackbar 네임스페이스 생성
kubect create namespace snackbar

# Service, Deployment 배포 (service.yaml)
kubectl apply -f service.yaml

# snackbar 네임스페이스의 모든 리소스 조회
kubectl get all -n snackbar

# order 컨테이너 환경변수 확인 
kubectl exec <order-pod> -n snackbar -- env | grep PAYMENT

# payment 컨테이너 환경변수 확인
kubectl exec <order-pod> -n snackbar -- env | grep ORDER

# snackbar 네임스페이스의 order Pod의 컨테이너 쉘 접속
kubectl exec -it <order-pod> -n snackbar -- sh

# payment 서비스 환경변수를 이용하여 Payment 호출
curl $PAYMENT_SERVICE_HOST:$PAYMENT_SERVICE_PORT

# payment 서비스의 로드밸런싱 확인 
for i in `seq 1 10`; do curl -s $PAYMENT_SERVICE_HOST:$PAYMENT_SERVICE_PORT; done
```
<br/><br/>
# 5. ClusterIP 서비스로 Pod 노출하는 방법(DSN 서버)
- 쿠버네티스에서 사용하는 도메인 이름 규칙 - FQDN(Fully Qualified Domain Name)
- /etc/resolv.conf 파일의 접미사를 통해 service 이름으로 팟을 호출, 응답 화인할 수 있다.

(45분 진행중) https://fastcampus.co.kr/courses/213060/clips/
<br/><br/><br/>
