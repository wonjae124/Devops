# 0.Sentiment analysis, Tokenizer, RNN, LSTM 기초 실습

# 1.주요 관점 정리

    - 최신 모델 아키텍쳐(Bert, Transfomer)보다 데이터셋 선정이 중요하다고 생각함.
    - BERT 모델은 기존에 영문 기사, 신문, 트위터의 데이터로 sentiment analysis에 학습된 사전모델이 있다.
    그러나, 이러한 데이터는 HDFS 로그 데이터와 연관이 없어서 적합한 수단이 아니다.
    따라서, 처음부터 학습 하기에 적합한 기존 Bi-LSTM 모델로 맥락을 학습이 적합하다고 판단했다.
    - Tokenizer를 통해, 문장을 최대 길이로 자른다. 
    - 딥러닝 모델은 시퀀스를 임베딩 레이어로 받아들여, 고정 길이로 벡터화 한다. 이에 의해, 앞뒤 맥락을 고려하게 된다. 
    - train, validation, test 데이터셋 비율은 8:1:1
        -6352개, 794개, 794개이다
    - 이벤트의 개수는 총 19개
    - 모델은 
    - F1 score 
    - Hyper parameter : batch_size =32, epochs = 50, earlystopping, LR_scheduler(step_decay, start : 0.01)
    - 

# 2.주요 기능 구현 코드

# 3.예측결과


```
$ sudo apt install postgresql postgresql-contrib 

$ sudo systemctl start postgresql.service

$ go mod init tutorial 

$ sudo -u postgres createuser --interactive

$ sudo -u postgres psql 

$ sudo adduser wonjae

$ sudo -i -u wonjae psql

```


### 출처

- 파이썬 텍스트 마이닝 완벽 가이드
