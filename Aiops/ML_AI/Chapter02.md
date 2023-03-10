# 0.Sentiment analysis, Tokenizer, RNN, LSTM 기초 실습

# 1.주요 관점 정리
    - BERT 모델은 기존에 영문 기사, 신문, 트위터의 데이터로 sentiment analysis에 학습된 사전모델이 있다.
    그러나, 이러한 데이터는 HDFS 로그 데이터와 연관이 없어서 적합한 수단이 아니다.
    따라서, 처음부터 학을 하기에 적합한 기존 Bi-LSTM 모델로 맥락을 학습이 적합하다고 판단했다.



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
