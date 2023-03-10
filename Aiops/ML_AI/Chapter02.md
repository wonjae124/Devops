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
    - HDFS데이터셋은 24시간 안의 데이터이며. 내가 가진 데이터의 Time 확인 결과, 203518부터 121018으로 약 4시간의 기록이다.

# 2.주요 기능 구현 코드
logparser 이용해서 다른 데이터 HDFS 데이터셋으로부터 다양한 로그를 받아들임.
`sudo docker run --name logparser -v /home/won/바탕화면/go_test/megazone/nlp/practice:/megazone -it e39c9d4c10d9` 를 이용해서 폴더 공유
기존 데이터셋의 데이터는 81109로 2008년 11월 09일이며, logparser를 통해 얻은 데이터는 대부분 081110으로 10일 즉, 다음날의 데이터이다.
- Drain demo.py의 파싱 저장 위치를 도커에서 공유한 로컬 폴더의 데이터셋 위치에 저장하도록 변경함 
<img src = "https://github.com/wonjae124/Devops/blob/main/image/%EC%8A%A4%ED%81%AC%EB%A6%B0%EC%83%B7%202023-03-10%2013-33-59.png?raw=true" width = 800>


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
