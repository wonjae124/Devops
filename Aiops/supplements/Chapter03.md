# 0. metric 
kubelet으로부터 리소스 매트릭을 수집
이를, k8s api 서버에 Metrics Api 형태로 노출된다.
Metrics api는 자원의 사용량에 따라, pod의 개수 및 용량을 제한하는 오토 스케일링을 기능을 한다.



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
 

# 2. 



#### 출처
- [Youtube - Webinar: Log Analysis with Machine Learning to Find Root Cause Faster](https://youtu.be/MpYB4Qcl570 )
- [slow-news](https://slownews.kr/86121)
