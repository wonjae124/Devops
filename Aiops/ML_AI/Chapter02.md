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

# 2. Issue 정리
- 추가적으로, logparser 이용해서 다른 데이터셋인 HDFS_2k로부터 다양한 로그를 받아옴.
       `sudo docker run --name logparser -v /home/won/바탕화면/go_test/megazone/nlp/practice:/megazone -it e39c9d4c10d9` 를 이용해서 폴더 공유

- 기존 데이터셋의 데이터는 81109로 2008년 11월 09일이며, logparser를 통해 얻은 데이터는 대부분 081110으로 10일 즉, 다음날의 데이터이다.

- Drain demo.py의 파싱 저장 위치를 도커에서 공유한 로컬 폴더의 데이터셋 위치에 저장하도록 변경함 
    <img src = "https://github.com/wonjae124/Devops/blob/main/image/%EC%8A%A4%ED%81%AC%EB%A6%B0%EC%83%B7%202023-03-10%2013-33-59.png?raw=true" width = 800>
- EventID의 제대로 regrex 적용 안됨. Issue 발생 -> Parsing은 생략하고, 패키지에서 제대로 parsing 해준 데이터셋을 대안으로 사용
    <img src = "https://github.com/wonjae124/Devops/blob/main/image/%EC%8A%A4%ED%81%AC%EB%A6%B0%EC%83%B7%202023-03-10%2015-49-30.png" width=200>
<br/><br/>

# 3. 모델 학습 & 기능 구현
- [Source code](https://github.com/wonjae124/Devops/blob/main/Aiops/ML_AI/sentiment_analysis_rev.ipynb)
<br/><br/>

### 출처

- 파이썬 텍스트 마이닝 완벽 가이드
<br/><br/><br/>
