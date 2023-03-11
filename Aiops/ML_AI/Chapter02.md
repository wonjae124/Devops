# 0. Sentiment analysis, Tokenizer, RNN, LSTM

# 1. 주요 관점 정리
- 최신 모델 아키텍쳐 외에 데이터셋 선정, 분석, 기능 구현에 집중함.
- BERT 모델은 기존에 영문 기사, 신문, 트위터의 데이터로 sentiment analysis에 학습된 사전모델이 있다.
    그러나, 이러한 데이터는 HDFS 로그 데이터와 연관이 없어서 적합한 수단이 아니다.
    따라서, 처음부터 학습 하기에 적합한 기존 Bi-LSTM 모델로 맥락 학습이 적합하다고 판단했다.
- Tokenizer를 통해, 문장을 최대 길이로 자른다. 
- 딥러닝 모델은 시퀀스를 임베딩 레이어로 받아들여, 고정 길이로 벡터화 한다. 이에 의해, 앞뒤 맥락을 고려하게 된다. 
- train, validation, test 데이터셋 비율은 8:1:1로 각각 6352개, 794개, 794개이다
- 모델은 기본적인 임베딩, RNN, Bi-LSTM으로 성능 비교, Test set에서 전부 성능 동일하게 나타남.
    - lrscheduler,early stop으로 인한 하이퍼 파라미터 튜닝 결과로 최적화 되었다고 생각함 
- Hyper parameter : batch_size =32, epochs = 50, earlystopping, LR_scheduler(step_decay, start : 0.01)
- HDFS_100k 데이터의 Time 확인 결과, 20:35:18부터 12:10:18으로 약 3시간 20분의 로그임
<br/><br/>

# 2. Issue 정리
- 추가적으로, logparser 이용해서 다른 데이터셋인 HDFS_2k로부터 다양한 로그를 받아옴.
    - `sudo docker run --name logparser -v /home/won/바탕화면/go_test/megazone/nlp/practice:/megazone -it e39c9d4c10d9` 를 이용해서 폴더 공유

- 기존 데이터셋의 데이터는 81109로 2008년 11월 09일이며, logparser를 통해 얻은 데이터는 대부분 081110으로 10일 즉, 다음날의 데이터이다.

- Drain demo.py의 파싱 저장 위치를 도커에서 공유한 로컬 폴더의 데이터셋 위치에 저장하도록 변경함 
    <img src = "https://github.com/wonjae124/Devops/blob/main/image/%EC%8A%A4%ED%81%AC%EB%A6%B0%EC%83%B7%202023-03-10%2013-33-59.png?raw=true" width = 800>
- EventID의 제대로 regrex 적용 안됨. Issue 발생 -> Parsing은 생략하고, 패키지에서 제대로 parsing 해준 데이터셋을 대안으로 사용
    <img src = "https://github.com/wonjae124/Devops/blob/main/image/%EC%8A%A4%ED%81%AC%EB%A6%B0%EC%83%B7%202023-03-10%2015-49-30.png" width=200>
    - ubuntu22.04에 최적화(librhash 업데이트)가 필요했던 것으로 추측함. 
    - `sudo apt install build-essential pkg-config cmake cmake-qt-gui ninja-build valgrind`
- GPU가 작동 안하고 있었음. 현재 RTX3050의 Cuda 11.6임. tensorflow는 최대 cuda 11.2를 지원함. 이에 11.6에서 11.2로 낮추어서 의존성 해결
    - 버전	파이썬 버전	컴파일러	빌드 도구	cuDNN	쿠다
    - 텐서플로우-2.11.0	3.7-3.10	GCC 9.3.1	바젤 5.3.0	8.1	11.2
<br/><br/>
    - pytorch와 tensorflow를 이용하는 사람이라면 nvcc version을 기준으로 설치하기(nvdia-smi 

# 3. 모델 학습 & 기능 구현
- [Source code](https://github.com/wonjae124/Devops/blob/main/Aiops/ML_AI/sentiment_analysis_rev.ipynb)
<br/><br/>
- 기존 로그는 이벤트가 다수였으나, 이벤트가 한 건인 새로운 로그로 예측한 결과, 부정이 긍정보다 압도적으로 많았다.
- 이처럼, 다수가 아닌 한 개의 이벤트에 대해서는 예측 성능이 낮다고 생각된다.

# 4. 배운점, 느낀점
- NLP가 사람처럼 맥락을 기억하려는 쪽으로 개선되고 있음을 이해했다.
- 로그를 parsing하는 여러 기법이 존재한다는걸 이해했다.
    - 이번에 이용했던 logparser 패키지의 regrex가 비정상 작동으로 인해, 잦은 유지보수가 이루어지는 commit을 보니, 굉장히 답답할 것 같다는 생각이 들었다. 이에, 기여하고 싶다는 생각이 들었다.
- 졸업작품인 데이터셋 융합 프로젝트의 개인 github 코드를 참고하여, 단기간에 사전과제를 진행했다.
    - 그 때, 진행하지 못했던 새로운 데이터에 대한 예측을 이번에는 진행하였기에, 보람찬 시간이었다.
    - 텐서플로우로 GPU 훈련, 예측을 해보는 좋은 기회였다
    
# 5. 아쉬운 점
- 하둡이나 아파치 서버에 대한 도메인 지식이 없어, 직접 모니터링 서버 로그를 받아 진행하지 못한 점이 아쉬웠다. 이후에는 실시간 클라우드의 트래픽이나 로그를 프로메테우스로 모니터링 해서 로그 결과를 저장하는 프로젝트를 진행해보고 싶다. 이는 사전과제에서 진행하지 못하여 아쉬웠는데, 이후 스터디를 꾸려서라도 진행 해야겠다.
- 맥락을 더 잘 이해하는 transfomer 기법은 이번에 도입하지 못해 아쉬웠다. 다음에는 trasfomer로 실습을 해봐야겠다. 
- pytorch2.0이 올 1월에 출시되었는데, 한 번 사용해보고 싶다

### 출처

- 파이썬 텍스트 마이닝 완벽 가이드
- [github - Ubuntu 22.04 for Deep Learning](https://gist.github.com/amir-saniyan/b3d8e06145a8569c0d0e030af6d60bea)
- [tensorflow cuda check](https://www.tensorflow.org/install/source?hl=ko#gpu)

<br/><br/><br/>

