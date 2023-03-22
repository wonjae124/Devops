# 0. 모두의 MLOPS
## 목표

- [x] minikube, CSI 구축, 클러스터 생성
- [x] NVIDIA-Docker를 Default Container Runtime으로 설정 
     - 쿠버네티스는 기본적으로 Docker-CE를 Default Container Runtime으로 사용합니다. 
        따라서, Docker Container 내에서 NVIDIA GPU를 사용하기 위해서는 NVIDIA-Docker 를 Container Runtime 으로 사용하여 pod를 생성할 수 있도록
        Default Runtime을 수정해 주어야 합니다
- [ ] MLFlow componenet
- [ ] API Deployment 

# 1. information
- Kustomize란
    - kustomize는 기존 쿠버네티스 리소스 정의(yaml파일)을 변경하지 않고 필드를 재정의하여 새로운 쿠버네티스 리소스를 생성하는 도구입니다.
    - 필드를 재정의 하는 것 때문에 kustomize가 재사용성이 좋다고 표현하기도 합니다.
    - kustomize는 helm과 많이 비교대상으로 언급이 되고 있는데요. helm은 템플릿에 정의된 필드의 값만 수정할 수 있습니다. 
      하지만 kustomize는 새로운 필드를 삽입할 수도 있고 기존 필드 값을 변경할 수 있습니다.그리고 helm은 버전(rivision)관리가 되지만 kustomize는 버전관리가 없습니다. 

- Helm이란?
    - Helm은 brew, apt-get 와 yaml 같은 저장소를 가지고있다.
    - 자동화가 잘 되어 있다.
    - go template을 지원한다.

- Seldon-Core란?
    - Seldon-Core는 쿠버네티스 환경에 수많은 머신러닝 모델을 배포하고 관리할 수 있는 오픈소스 프레임워크 중 하나입니다.

- 프로메테우스&그라파나란?
    - 모니터링을 위한 도구입니다.안정적인 서비스 운영을 위해서는 서비스와 서비스가 운영되고 있는 인프라의 상태를 지속해서 관찰하고, 관찰한 메트릭을 바탕으로 문제가 생길 때 빠르게 대응해야 합니다.
    - 쿠버네티스 클러스터에 프로메테우스와 그라파나를 설치한 뒤, Seldon-Core 로 생성한 SeldonDeployment 로 API 요청을 보내, 정상적으로 Metrics 이 수집되는지 확인

# 2. model API deployment
- 실제 서비스에서 머신러닝이 사용될 때는 API를 이용해서 학습된 모델을 사용합니다. 
  모델은 API 서버가 구동되는 환경에서 한 번만 로드가 되며, DNS를 활용하여 외부에서도 쉽게 추론 결과를 받을 수 있고 다른 서비스와 연동할 수 있습니다.

- 모델을 API로 만드는 작업에는 생각보다 많은 부수적인 작업이 필요합니다.
  그래서 API로 만드는 작업을 더 쉽게 하기 위해서 Tensorflow와 같은 머신러닝 프레임워크 진영에서는 추론 엔진(Inference engine)을 개발하였습니다.
  추론 엔진들을 이용하면 해당 머신러닝 프레임워크로 개발되고 학습된 모델을 불러와 추론이 가능한 API(REST 또는 gRPC)를 생성합니다.
  이러한 추론 엔진을 활용하여 구축한 API 서버로 추론하고자 하는 데이터를 담아 요청을 보내면, 추론 엔진이 추론 결과를 응답에 담아 전송하는 것입니다.

  Tensorflow : Tensorflow Serving
  PyTorch : Torchserve
  Onnx : Onnx Runtime
  오프소스에서 공식적으로 지원하지는 않지만, 많이 쓰이는 sklearn, xgboost 프레임워크를 위한 추론 엔진도 개발되어 있습니다.

  이처럼 모델의 추론 결과를 API의 형태로 받아볼 수 있도록 배포하는 것을 API Deployment라고 합니다.

reference
- [모두의 MLops](https://mlops-for-all.github.io/docs/setup-kubernetes/setup-nvidia-gpu/)
